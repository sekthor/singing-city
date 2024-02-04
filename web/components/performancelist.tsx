import { Performance } from "@/model/performance";

type PerformanceListProps = {
    performances: Performance[];
}

export default function PerformanceList({performances}:PerformanceListProps ) {
    return (
        <>
            <p>This is a list of performances</p>
        </>
    )
}