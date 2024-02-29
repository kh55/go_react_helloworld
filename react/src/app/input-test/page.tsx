"use client";

import axios from "axios";
import { useState } from "react";

import styles from './input-test.module.css'


export default function Page() {
    const [message, setMessage] = useState('wait');

    async function action_test(form) {
        try {
            const message = form.get('message')
            const url = "http://localhost:3001/input";
            const data = {
                InputData: message,
            };
            const response = await axios.post(url, data);
            console.log(response.data);
            setMessage(response.data['message']);
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    }

  return (
    <>
      <h1>Server Action Test</h1>
      <form action={action_test}>
        <input className={styles.inputbox} type="text" name="message" />
        <button type="submit">OK</button>
      </form>
      <p>{message}</p>
    </>
  )
}
