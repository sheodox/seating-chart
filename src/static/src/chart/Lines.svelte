<style>
	.chart {
		position: relative;
		top: 0;
		width: 100%;
		height: 100%;
	}
	.chart.editable {
		cursor: crosshair;
	}
	svg {
		pointer-events: none;
		stroke-width: 2px;
		stroke: white;
		fill: none;
	}
	.point-button {
		margin: 0;
		background: white;
		color: gray;
		position: absolute;
		transform: translate(-100%, -50%);
		line-height: 1;
		font-size: 2rem;
		padding: 0;
		cursor: move;
	}
	.point-button.moving {
		pointer-events: none;
	}
</style>

<svelte:window on:resize={updateCanvasDimensions} />
<div
	class="chart"
	class:editable
	on:click={addPoint}
	bind:this={canvas}
	on:drop={finishMovingPoint}
	on:mouseup={finishMovingPoint}
	on:mousemove={movePoint}
	on:dragover|preventDefault
>
	<svg viewBox="0 0 {canvasRes} {canvasRes}" xmlns="http://www.w3.org/2000/svg">
		<path d={path} />
	</svg>

	{#if editable}
		{#each coordinates as coord, index}
			{#if index !== movingPoint}
				<button
					class:moving={movingPoint !== null}
					class="point-button"
					draggable="true"
					on:dragstart={(e) => moveStart(e, index)}
					on:contextmenu|preventDefault={() => deletePoint(index)}
					on:dblclick={() => duplicatePoint(index)}
					on:click|stopPropagation
					style="top: {coord.y * 100}%; left: {coord.x * 100}%">{index}</button
				>
			{/if}
		{/each}
	{/if}
</div>

<script lang="ts">
	import { onMount, tick } from 'svelte';
	import { editorMode } from '../stores/editor';

	$: editable = $editorMode === 'lines';
	type Coordinate = { x: number; y: number };

	const canvasRes = 1000;

	let canvas: HTMLElement,
		canvasDimensions = { height: 0, width: 0, top: 0 },
		coordinates: Coordinate[] = [],
		movingPoint: number = null;

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

	onMount(() => {
		updateCanvasDimensions();
	});

	function addPoint(e: MouseEvent) {
		if (editable) {
			coordinates = [...coordinates, genCoordinatesFromEvent(e.offsetX, e.offsetY)];
		}
	}

	function genCoordinatesFromEvent(x: number, y: number) {
		return {
			x: x / canvasDimensions.width,
			y: y / canvasDimensions.height,
		};
	}

	$: path = genPath(coordinates);

	function genPath(coordinates: Coordinate[]) {
		if (coordinates.length === 0) {
			return '';
		}
		const start = coordinates[0],
			points = coordinates
				.slice(1)
				.map((coord) => {
					return `L${coord.x * canvasRes},${coord.y * canvasRes}`;
				})
				.join(' ');

		return `M${start.x * canvasRes},${start.y * canvasRes} ${points} z`;
	}

	function deletePoint(index: number) {
		if (editable) {
			coordinates.splice(index, 1);
			coordinates = coordinates;
		}
	}

	function moveStart(e: DragEvent, index: number) {
		movingPoint = index;
		e.dataTransfer.dropEffect = 'move';
	}

	function finishMovingPoint(e: MouseEvent) {
		if (movingPoint === null) {
			return;
		}

		movePoint(e);
		movingPoint = null;
	}

	function movePoint(e: MouseEvent) {
		if (movingPoint === null) {
			return;
		}

		coordinates[movingPoint] = genCoordinatesFromEvent(e.offsetX, e.offsetY);
		coordinates = coordinates;
	}

	function duplicatePoint(index: number) {
		const newPoint = {
			x: coordinates[index].x + 10,
			y: coordinates[index].y + 10,
		};
		coordinates = [...coordinates.slice(0, index), newPoint, ...coordinates.slice(index)];
	}
</script>
