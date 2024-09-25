import React, { useState } from 'react';

const URLShortener = () => {
  const [url, setUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [error, setError] = useState('');
  const baseURL = typeof window !== "undefined" ? window.location.origin + "/" : "http://localhost:8080/";
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError('');
    setShortUrl('');

    try {
      const response = await fetch('http://localhost:8080/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to create short URL');
      }

      const data = await response.json();
      setShortUrl(data.shortCode);
    } catch (err: any) {
      setError(err.message);
    }
  };

  const handleCopy = () => {
    navigator.clipboard.writeText(baseURL + shortUrl)
      .then(() => {
        setError('Short URL copied to clipboard!');
      })
      .catch(() => {
        setError('Failed to copy short URL');
      });
  };

  return (
    <div className="max-w-md mx-auto p-4">
      <form onSubmit={handleSubmit} className="flex flex-col">
        <input
          type="url"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="Enter your URL"
          required
          className="border border-gray-300 p-2 rounded mb-4 focus:outline-none focus:ring focus:ring-blue-400"
        />
        <button
          type="submit"
          className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600 transition duration-200"
        >
          Shorten URL
        </button>
      </form>
      {shortUrl && (
        <div className="mt-4 p-2 rounded">
          {error && (
            <div className="mt-4 text-green-500 text-center">
              <p>{error}</p>
            </div>
          )}
          <div className='flex justify-center items-center mt-5'>
            <button
              onClick={handleCopy}
              className="mt-2 bg-green-500 text-white font-bold p-2 mt-5 rounded hover:bg-green-600 transition duration-200 bg-blue-500"
            >
              Copy to Clipboard
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default URLShortener;

