import About from "@/components/home/about/About";
import Navbar from "@/components/home/navbar/navbar";
import HowWeWork from "@/components/home/work/work";

export default function Home() {
  return (
    <div className="">
      <div className=" h-24">
        <Navbar />
      </div>
      <About />
      <HowWeWork />
    </div>
  );
}
