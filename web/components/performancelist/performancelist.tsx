"use client"

import { Performance } from "@/model/performance";
import { Table } from "@/components/ui/table";

type PerformanceListProps = {
    performances: Performance[];
}

export default function PerformanceList({performances}:PerformanceListProps ) {
    return (
        <>
            <h1>Performances</h1>

            <Table>

            </Table>
        </>
    )
}