<style lang="scss">
	.tables {
		position: absolute;
		top: 0;
		width: 100%;
		height: 100%;
		&:not(.editable) {
			opacity: 0.2;
			pointer-events: none;
		}
		&.mode-tables {
			cursor: crosshair;

			div.table {
				cursor: move;
			}
		}
		&.mode-assignments {
			.table:hover {
				border-radius: 5px;
				overflow: auto;
				.guests {
					display: block;
				}
			}
		}
	}
	.table {
		background: var(--shdx-gray-50);
		outline: 4px solid transparent;
		width: 10%;
		aspect-ratio: 1;
		position: absolute;
		border-radius: 50%;
		transform: translate(-50%, -50%);
		color: black;
		display: flex;
		align-items: center;
		border: 1px solid black;
		transition: border-radius 0.1s;
		overflow: hidden;

		.guests {
			font-size: var(--shdx-font-size-1);
		}

		&.moving {
			pointer-events: none;
		}

		.table-name {
			color: black;
			border-radius: 2px;
			padding: 2px 3px;
			font-weight: bold;
			&:hover {
				background: var(--shdx-gray-200);
			}

			&.duplicate-name {
				border: 2px solid red;
				border-radius: 5px;
				background: var(--shdx-red-800);
				color: red;
			}
		}
	}
	.table.highlight-table {
		outline-offset: 4px;
		outline-color: var(--shdx-blue-500);
	}

	.dragging-guest {
		outline-offset: 0px;

		&.can-hold {
			outline-color: var(--shdx-blue-500);
		}
		&:not(.can-hold) {
			outline-color: var(--shdx-red-500);
			opacity: 0.3;
		}
	}
	.over-capacity {
		background: var(--shdx-red-400);
	}
</style>

<div
	class="tables mode-{$editorMode}"
	class:editable
	on:click={canvasClick}
	bind:this={canvas}
	on:drop={finishMovingPoint}
	on:mouseup={finishMovingPoint}
	on:dragover|preventDefault
>
	{#each $tables as table, index}
		<div
			class="table f-column"
			style="top: {table.posY * 100}%; left: {table.posX * 100}%;"
			class:moving={movingTable !== null}
			class:dragging-guest={!!$draggingGuest}
			class:highlight-table={$highlightTable === table.id}
			class:can-hold={$draggingGuest && canHoldGuest(table, $guests)}
			class:over-capacity={peopleAtTable(table, $guests) > table.capacity}
			draggable={$editorMode === 'tables'}
			on:dragstart={(e) => moveStart(e, index)}
			on:drop={(e) => assignmentDrop(e, table)}
			on:contextmenu|preventDefault={() => deleteTable(index)}
			on:click|stopPropagation
			on:dragover|preventDefault={() => dragOver(table)}
			on:dragleave={() => ($highlightTable = '')}
		>
			<button
				class="table-name"
				class:duplicate-name={hasDuplicateName(table.name)}
				title={hasDuplicateName(table.name) ? 'Duplicate name' : ''}
				on:click|stopPropagation={() => renameTable(index)}>{table.name}</button
			>
			<p class="fw-bold m-0">
				<Icon icon="users" />
				{peopleAtTable(table, $guests)} / {table.capacity}
			</p>
			<div class="guests">
				{#each assignedGuests(table, $guests) as guest}
					<p class="m-0">
						{guest.firstName}
						{guest.lastName}
						- ({guest.people})
					</p>
				{/each}
			</div>
		</div>
	{/each}
</div>
<svelte:window on:resize={updateCanvasDimensions} />

<script lang="ts">
	import { onMount, tick } from 'svelte';
	import { Icon } from 'sheodox-ui';
	import { editorMode } from '../stores/editor';
	import { guestOps, guests, Guest, draggingGuest } from '../stores/guests';
	import { tableOps, tables, highlightTable, Table } from '../stores/tables';

	$: editable = ['tables', 'assignments'].includes($editorMode);

	let canvas: HTMLElement,
		canvasDimensions = { height: 0, width: 0, top: 0 },
		movingTable: number = null,
		movingTableShift = {
			x: 0,
			y: 0,
		};

	function peopleAtTable(table: Table, guests: Guest[], notIncluding?: Guest) {
		return guests.reduce((people, guest) => {
			if (!guest.going || notIncluding === guest) {
				return people;
			}
			return people + (guest.tableId === table.id ? guest.people : 0);
		}, 0);
	}

	function canHoldGuest(table: Table, guests: Guest[]) {
		return peopleAtTable(table, guests, $draggingGuest) + ($draggingGuest?.people || 0) <= table.capacity;
	}

	function assignedGuests(table: Table, guests: Guest[]) {
		return guests.filter((g) => g.going && g.tableId === table.id);
	}

	function canvasClick(e: MouseEvent) {
		if ($editorMode !== 'tables') {
			return;
		}

		const { posX, posY } = genCoordinatesFromEvent(e.offsetX, e.offsetY);

		tableOps.add({
			name: `Table ${$tables.length + 1}`,
			posX,
			posY,
		});
	}

	function dragOver(table: Table) {
		$highlightTable = table.id;
	}

	$: $editorMode && updateCanvasDimensions();
	$: canvas && updateCanvasDimensions();

	async function updateCanvasDimensions() {
		await tick();
		if (canvas) {
			canvasDimensions = {
				height: canvas.offsetHeight,
				width: canvas.offsetWidth,
				top: canvas.offsetTop,
			};
		}
	}

	function hasDuplicateName(name: string) {
		return $tables.filter((table) => table.name === name).length > 1;
	}

	function assignmentDrop(e: DragEvent, table: Table) {
		guestOps.assign(e.dataTransfer.getData('guestId'), table.id);
		$highlightTable = '';
	}

	function deleteTable(index: number) {
		if (editable && confirm('Are you sure you want to delete this table?')) {
			tableOps.delete($tables[index].id);
		}
	}

	function renameTable(index: number) {
		const newName = prompt('Enter a new name', $tables[index].name)?.trim();
		if (newName) {
			const newTable = tableOps.edit({
				...$tables[index],
				name: newName,
			});
		}
	}

	function moveStart(e: DragEvent, index: number) {
		movingTable = index;
		const { left, top, width, height } = (e.target as HTMLElement).getBoundingClientRect();
		movingTableShift = {
			x: (e.clientX - (left + width / 2)) / canvasDimensions.width,
			y: (e.clientY - (top + height / 2)) / canvasDimensions.width,
		};
		e.dataTransfer.dropEffect = 'move';
	}

	function finishMovingPoint(e: MouseEvent) {
		if (movingTable === null) {
			return;
		}

		moveTable(e);
		movingTable = null;
		movingTableShift = { x: 0, y: 0 };
	}

	function moveTable(e: MouseEvent) {
		if (movingTable === null) {
			return;
		}

		const newTable = {
			...$tables[movingTable],
			...genCoordinatesFromEvent(e.offsetX, e.offsetY),
		};

		tableOps.edit(newTable);
	}

	onMount(() => {
		updateCanvasDimensions();
	});

	function genCoordinatesFromEvent(x: number, y: number) {
		return {
			posX: x / canvasDimensions.width - movingTableShift.x,
			posY: y / canvasDimensions.height - movingTableShift.y,
		};
	}
</script>
