import type { MetaFunction } from "@remix-run/node";
import URLShortener from "~/components/URLShorter";

export const meta: MetaFunction = () => {
  return [
    { title: "New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function Index() {
  return (
    <>
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <div className="bg-white shadow-md rounded-lg p-8 max-w-md w-full">
          <h1 className="text-5xl font-bold text-center mb-6 pixel">URL Shortener</h1>
          <URLShortener />
        </div>
      </div>
    </>
  );
}

