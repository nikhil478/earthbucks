import type { MetaFunction } from '@remix-run/node'

export const meta: MetaFunction = () => {
  return [
    { title: "SatoRed (Satoshi's Redemption)" },
    { name: 'description', content: 'Welcome to SatoRed!' },
  ]
}

export default function Index() {
  return (
    <div className="">
      <div className="mb-4 mt-4 flex">
        <div className="mx-auto">
          <div className="inline-block align-middle">
            <img
              src="/satored-logo.png"
              alt=""
              className="mx-auto mb-4 block h-[200px] w-[200px]"
            />
            <div className="hidden dark:block">
              <img
                src="/satored-text-white.png"
                alt="SatoRed"
                className="mx-auto block h-[50px]"
              />
            </div>
            <div className="block dark:hidden">
              <img
                src="/satored-text-black.png"
                alt="SatoRed"
                className="mx-auto block h-[50px]"
              />
            </div>
            <div className="mb-4 text-center text-black dark:text-white">
              42 million ninjas. No pre-mine. GPUs. Big blocks. Script.
              <br />
              <br />
              Genesis block in 2024.
              <br />
              <br />
              <a
                href="https://t.me/+e-PoBNNgdDU3MDUx"
                target="_blank"
                className="underline"
              >
                Telegram
              </a>{' '}
              &middot;{' '}
              <a
                href="https://discord.gg/6hyVpAVPC5"
                target="_blank"
                className="underline"
              >
                Discord
              </a>{' '}
              &middot;{' '}
              <a
                href="https://x.com/satored_com"
                target="_blank"
                className="underline"
              >
                X (Twitter)
              </a>{' '}
              &middot;{' '}
              <a
                href="https://github.com/satored/satored"
                target="_blank"
                className="underline"
              >
                GitHub
              </a>{' '}
              &middot;{' '}
              <a
                href="https://reddit.com/r/satored"
                target="_blank"
                className="underline"
              >
                reddit
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
