import { NextPage } from 'next'
import Head from 'next/head'

type SSGProps = {}

const SSG: NextPage<SSGProps> = () => {
  return (
    <div>
      <Head>
        <title>Static Site Generation</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <p>
          このページは静的サイト生成によって生成されています。
        </p>
      </main>
    </div>
  )
}

export default SSG