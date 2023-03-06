import axios from 'axios'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { useState } from 'react'

type SSGProps = {}

const Login: NextPage<SSGProps> = () => {
  const router = useRouter()

  const [userId, setUserId] = useState('');
  const [password, setPassword] = useState('');

  const onClick = () => {
    const data = { userId, password }
    const url = 'http://localhost:8080/api/login'
    axios.post(url, data, { withCredentials: true })
      .then((res) => {
        console.log(res)
        router.push('/dashboard')
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
    <>
      <form action="" method="post">
        <table>
          <tbody>
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
          </tbody>
        </table>
      </form>
    </>
  )
}

export default Login