import style1 from './test1.module.css'
import style2 from './test2.module.css'

export default function Page() {
  return (
    <>
      <h1>CSS Test</h1>
      <div className={style1.test1}>Test1 message.</div>
      <div className={style2.test2}>Test2 message.</div>
    </>
  )
}