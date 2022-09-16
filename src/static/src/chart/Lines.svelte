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
		position: absolute;
		top: 0;
		pointer-events: none;
		stroke-width: 2px;
		stroke: white;
		fill: none;
	}
	svg.inactive {
		opacity: 0.3;
	}
	svg.highlighted {
		stroke: var(--shdx-blue-500);
		opacity: 1;
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

<div
	class="chart"
	class:editable={editable && editingLine}
	on:click={addPoint}
	bind:this={canvas}
	on:drop={finishMovingPoint}
	on:mouseup={finishMovingPoint}
	on:mousemove={movePoint}
	on:dragover|preventDefault
>
	{#each $lines as line, lineIndex}
		<svg
			viewBox="0 0 {canvasRes} {canvasRes}"
			xmlns="http://www.w3.org/2000/svg"
			class:highlighted={$highlightingLineIndex === lineIndex}
			class:inactive={editable && $editingLineIndex !== lineIndex}
		>
			<path d={genPath(line)} />
		</svg>
	{/each}

	{#if editable && editingLine}
		{#each editingLine.coords as coord, index}
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
	import { editingLineIndex, lines, lineOps, highlightingLineIndex } from '../stores/lines';
	import type { Line } from '../stores/lines';
	import { editorMode } from '../stores/editor';

	$: editable = $editorMode === 'lines';

	const canvasRes = 1000;

	let canvas: HTMLElement,
		movingPoint: number = null;

	$: editingLine = $lines[$editingLineIndex];

	function addPoint(e: MouseEvent) {
		if (editable && editingLine) {
			const updatedLine = cloneLine(editingLine);
			updatedLine.coords.push(genCoordinatesFromEvent(e.offsetX, e.offsetY));
			lineOps.edit(updatedLine);
		}
	}

	function genCoordinatesFromEvent(x: number, y: number) {
		return {
			x: x / canvas.offsetWidth,
			y: y / canvas.offsetHeight,
		};
	}

	function cloneLine(line: Line): Line {
		return {
			...line,
			coords: line.coords.map((coord) => ({ ...coord })),
		};
	}

	function genPath(line: Line) {
		if (line.coords.length === 0) {
			return '';
		}
		const start = line.coords[0],
			points = line.coords
				.slice(1)
				.map((coord) => {
					return `L${coord.x * canvasRes},${coord.y * canvasRes}`;
				})
				.join(' ');

		return `M${start.x * canvasRes},${start.y * canvasRes} ${points} ${line.closed ? 'z' : ''}`;
	}

	function deletePoint(index: number) {
		if (editable) {
			const updatedLine = cloneLine(editingLine);
			updatedLine.coords.splice(index, 1);
			lineOps.edit(updatedLine);
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
		lineOps.edit(editingLine);
		movingPoint = null;
	}

	function movePoint(e: MouseEvent) {
		if (movingPoint === null) {
			return;
		}

		editingLine.coords[movingPoint] = genCoordinatesFromEvent(e.offsetX, e.offsetY);
		editingLine = editingLine;
	}

	function duplicatePoint(index: number) {
		const updatedLine = cloneLine(editingLine);

		const newPoint = {
			x: updatedLine.coords[index].x + 0.02,
			y: updatedLine.coords[index].y + 0.02,
		};
		editingLine.coords = [...updatedLine.coords.slice(0, index), newPoint, ...updatedLine.coords.slice(index)];
		editingLine = editingLine;
	}
</script>
