import { useEffect, useState } from "react";
import { useParams } from "@remix-run/react";

export default function Token() {
  const { token } = useParams<{ token: string }>();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchTokenData = async () => {
      try {
        const response = await fetch(`http://localhost:8080/${token}`);

        if (!response.ok) {
          throw new Error('Token not valid');
        }

        const data = await response.json();
        const shortUrl = data.url;

        if (shortUrl) {
          window.location.href = shortUrl;
        } else {
          throw new Error('Short URL not found');
        }
      } catch (err: any) {
        console.error(err);
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchTokenData();
  }, [token]);

  return (
    <div className="flex items-center justify-center h-screen pixel text-5xl">
      {loading ? (
        <div className="flex items-center">
          <p className="mt-4">Loading your URL...</p>
        </div>
      ) : error ? (
        <h1 className="text-red-500">{error}</h1>
      ) : (
        <p>Redirecting...</p>
      )}
    </div>
  );
}

