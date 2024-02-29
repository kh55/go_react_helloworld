"use client";

import axios from "axios";
import { useState } from "react";

export default function Home() {
    const [message, setMessage] = useState('wait');

  const getData = async () => {
    try {
      const response = await axios.get("http://localhost:3001/redis");
      setMessage(response.data['message']);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          onClick={getData}
        >
          getData
        </button>
      </div>
      <p>{message}</p>
    </main>
  );
}
