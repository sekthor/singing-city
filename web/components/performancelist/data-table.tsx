"use client"

import { ColumnDef } from "@tanstack/react-table"
import { Performance } from "@/model/performance"

export const columns: ColumnDef<Performance>[] = [
    {
      accessorKey: "Venue",
      header: "Venue",
    },
    {
      accessorKey: "start",
      header: "Start",
    },
  ]