import { VideoPlayer } from "@/components/VideoPlayer";
import { Suspense } from "react";
import { VideoDetail } from "./VideoDetails";
import { unstable_after as after } from "next/server";

const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

async function incrementViews(id: string): Promise<void> {
  await sleep(5000);
}

export default async function VideoPlayPage({
  params,
}: {
  params: { id: string };
}) {
  after(async () => {
    await incrementViews(params.id);
  });

  return (
    <div>
      <VideoPlayer />
      <Suspense fallback={<div>loading...</div>}>
        <VideoDetail id={params.id} />
      </Suspense>
      <p className="text-primary">1000 Views</p>
    </div>
  );
}
