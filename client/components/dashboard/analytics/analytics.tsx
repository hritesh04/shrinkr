"use client";
import {
  ArcElement,
  Chart,
  Legend,
  LineController,
  LineElement,
  LinearScale,
  PointElement,
  Title,
  Tooltip,
  TimeScale,
  CategoryScale,
  TooltipItem,
} from "chart.js";
import { Line } from "react-chartjs-2";
import "chartjs-adapter-dayjs-4/dist/chartjs-adapter-dayjs-4.esm";
import { animation } from "@/utils/chartAnimation";
export default function Analytics({ url }: { url: string }) {
  Chart.register(
    CategoryScale,
    LineController,
    LineElement,
    PointElement,
    LinearScale,
    Tooltip,
    TimeScale,
    Title
  );

  return (
    <div className="bg-[#ffffff] rounded-lg">
      <Line
        data={{
          labels: [
            "2024-04-01",
            "2024-04-02",
            "2024-04-03",
            "2024-04-04",
            "2024-04-05",
            "2024-04-06",
            "2024-04-07",
            "2024-04-08",
            "2024-04-09",
            "2024-04-10",
            "2024-04-11",
            "2024-04-12",
            "2024-04-13",
            "2024-04-14",
            "2024-04-15",
            "2024-04-16",
            "2024-04-17",
            "2024-04-18",
            "2024-04-19",
            "2024-04-20",
            "2024-04-21",
            "2024-04-22",
            "2024-04-23",
            "2024-04-24",
            "2024-04-25",
            "2024-04-26",
            "2024-04-27",
            "2024-04-28",
            "2024-04-29",
            "2024-04-30",
          ],
          datasets: [
            {
              label: url,
              data: [
                0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 3, 4, 5, 5, 6, 7, 7, 8, 10, 3,
                3, 4, 5, 5, 6, 7, 7, 8, 10,
              ],
              borderColor: "black", // color of line
              borderWidth: 1, // width of line
              pointRadius: 1, // radius of the dots
              pointHoverRadius: 3, // radius of dots at hover
              pointBackgroundColor: "black", // color of points on hover
            },
          ],
        }}
        options={{
          // @ts-ignore
          animation: animation,
          interaction: {
            mode: "index", // Hover over data points
            intersect: false, // Don't need line intersection for tooltip
          },
          plugins: {
            title: {
              display: true,
              color: "black",
              text: url,
            },
            tooltip: {
              mode: "index",
              intersect: false,
              boxPadding: 5,
              callbacks: {
                title: (context) => {
                  let date = new Date(context[0].parsed.x);
                  return date.toString().slice(0, 15);
                },
                label: (context) => {
                  const index = context.dataIndex;
                  const label = context.dataset.label;
                  const value = context.dataset.data[index];
                  return `${label}: ${value} Clicks`;
                },
                footer: (tooltipItems: TooltipItem<"line">[]) => {
                  let index = tooltipItems[0].dataIndex;
                  let totalClicks = 0;
                  for (let i = 0; i <= index; i++) {
                    totalClicks += Number(tooltipItems[0].dataset.data[i]);
                  }
                  return `Total Clicks : ${totalClicks}`;
                },
              },
            },
          },
          scales: {
            x: {
              ticks: {
                color: "black",
              },
              border: {
                color: "black",
              },
              type: "time",
              time: {
                displayFormats: {
                  month: "MM DD",
                },
                unit: "day",
              },
              grid: {
                display: false,
                tickWidth: 50,
              },
            },
            y: {
              ticks: {
                color: "black",
              },
              border: {
                color: "black",
              },
              title: {
                display: true,
                text: "Number of Clicks",
                color: "black",
              },
              grid: {
                display: false,
              },
            },
          },
        }}
      />
    </div>
  );
}
