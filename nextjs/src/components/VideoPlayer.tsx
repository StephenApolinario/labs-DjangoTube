"use client";
import { useEffect, useRef } from "react";

export function VideoPlayer() {
  const videoTagRef = useRef<HTMLVideoElement>(null);
  useEffect(() => {
    // Lib de player
    // Usar a ref
    // Start
  }, []);

  return <video controls autoPlay ref={videoTagRef}></video>;
}
