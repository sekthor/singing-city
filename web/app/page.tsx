import Image from "next/image";

import Navbar from "@/components/navbar"

export default function Home({ params }: { params: { lang: string }}) {
  return (
    <>
      <Navbar lang={params.lang} />
      <main>
      </main>
    </>
  );
}
