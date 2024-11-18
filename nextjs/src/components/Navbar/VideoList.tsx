import { Video } from "@/app/models";
import Link from "next/link";
import { VideoCard } from "../VideoCard";

const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

async function getVideos(): Promise<Video[]> {
  await sleep(5000);
  const cacheKey = "videos";

  try {
    const response = await fetch("http://localhost:8000/videos", {
      next: { revalidate: 60 }, // Cache por 60 segundos
      cache: "no-cache", // Tenta sempre retornar do cache primeiro
    });

    if (!response.ok) {
      throw new Error("Erro ao buscar vídeos");
    }

    const videos: Video[] = await response.json();

    return videos;
  } catch (error) {
    console.error("Erro ao buscar vídeos:", error);
    throw error; // Rejeita caso seja necessário tratamento adicional
  }
}

export async function VideoList() {
  const videos = await getVideos();

  return videos.map((video) => (
    <Link key={video.id} href={`video/${video.id}/play`}>
      <VideoCard
        title={video.title}
        thumbnail={video.thumbnail}
        views={video.num_views}
      />
    </Link>
  ));
}
