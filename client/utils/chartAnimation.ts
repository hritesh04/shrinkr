const totalDuration = 3000;
const previousY = (ctx: any) =>
  ctx.index === 0
    ? ctx.chart.scales.y.getPixelForValue(100)
    : ctx.chart
        .getDatasetMeta(ctx.datasetIndex)
        .data[ctx.index - 1].getProps(["y"], true).y;
export const animation = {
  x: {
    type: "number",
    easing: "linear",
    duration: 100,
    from: NaN,
    delay(ctx: any) {
      if (ctx.type !== "data" || ctx.xStarted) {
        return 0;
      }
      ctx.xStarted = true;
      return (
        (ctx.index * totalDuration) /
        ctx.chart.data.datasets[ctx.datasetIndex].data.length
      );
    },
  },
  y: {
    type: "number",
    easing: "linear",
    duration: 100,
    from: previousY,
    delay(ctx: any) {
      if (ctx.type !== "data" || ctx.yStarted) {
        return 0;
      }
      ctx.yStarted = true;
      return (
        ctx.index *
        (totalDuration / ctx.chart.data.datasets[ctx.datasetIndex].data.length)
      );
    },
  },
};
