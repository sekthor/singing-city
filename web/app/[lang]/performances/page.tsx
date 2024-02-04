import Navbar from "@/components/navbar";
import PerformanceList from "@/components/performancelist";
import { Performance } from "@/model/performance"

let performances: Performance[] = []

export default function PerformancePage() {
    return(
        <>
            <Navbar />
            <PerformanceList performances={performances} />
        </>
    )
}