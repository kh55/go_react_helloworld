"use client";

import axios from "axios";
import { useState } from "react";

import styles from '../../components/ui-elements/input-test.module.css'


export default function Page() {
    const [output_data, setMessage] = useState('wait');

    async function action_test(form) {
        try {
            const input_data = form.get('input_data')
            const url = "http://localhost:3001/input";
            const data = {
                InputData: input_data,
            };
            const response = await axios.post(url, data);
            console.log(response.data);
            setMessage(response.data['response_message']);
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    }

  return (
    <>
      <h1>送信テスト</h1>
      <form action={action_test}>
        <input className={styles.inputbox} type="text" name="input_data" />
        <button className={styles.sendbtn} type="submit">送信</button>
      </form>
      <p>入力した文字は：{output_data}</p>
    </>
  )
}
