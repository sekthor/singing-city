import Navbar from "@/components/navbar";
import PerformanceList from "@/components/performancelist/performancelist";
import { Performance } from "@/model/performance"

let performances: Performance[] = []

export default function PerformancePage({ params }: { params: { lang: string }}) {
    return(
        <>
            <Navbar lang={params.lang} />
            <PerformanceList performances={performances} />
        </>
    )
}