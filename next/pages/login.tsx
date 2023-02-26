import axios from 'axios'
import { NextPage } from 'next'
import Head from 'next/head'
import { useState } from 'react'

type SSGProps = {}

const Login: NextPage<SSGProps> = () => {
  const [userId, setUserId] = useState('');
  const [password, setPassword] = useState('');

  const onClick = () => {
    const data = { userId, password }
    const url = 'http://localhost:8080/api/login'
    axios.post(url, data)
      .then((res) => {
        console.log(res)
      })
      .catch((err) => {
        console.log(err)
      })
  }

  const handleUserIdChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUserId(() => e.target.value)
  }

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(() => e.target.value)
  }

  return (
    <div>
      <form action="" method="post">
        <table>
          <tr>
            <td>メールアドレス</td>
            <td><input type="text" value={userId} onChange={handleUserIdChange} required /></td>
          </tr>
          <tr>
            <td>パスワード</td>
            <td><input type="password" value={password} onChange={handlePasswordChange} required /></td>
          </tr>
          <tr>
            <td></td>
            <td><button type="button" onClick={onClick} >ログイン</button></td>
          </tr>
        </table>
      </form>
    </div>
  )
}

export default Login