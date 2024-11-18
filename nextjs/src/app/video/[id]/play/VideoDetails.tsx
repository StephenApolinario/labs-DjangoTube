import { Video } from "@/app/models";

async function getVideos(id: string): Promise<Video> {
  const cacheKey = "videos";

  try {
    const response = await fetch(`http://localhost:8000/videos/${id}`, {
      next: {
        tags: [`video:${id}`],
      },
      cache: "force-cache", // Tenta sempre retornar do cache primeiro
    });

    if (!response.ok) {
      throw new Error("Erro ao buscar vídeos");
    }

    const videos: Video = await response.json();

    return videos;
  } catch (error) {
    console.error("Erro ao buscar vídeos:", error);
    throw error; // Rejeita caso seja necessário tratamento adicional
  }
}

export async function VideoDetail({ id }: { id: string }) {
  const video = await getVideos(id);

  return (
    <div>
      <h1 className="text-primary"> Play do vídeo</h1>
      <p className="text-primary"> {id} </p>
      <p className="text-primary"> {video.title} </p>
    </div>
  );
}
