import * as React from "react"
import { Link } from "@tanstack/react-router"

export const Header = () => {
  return (
    <header className="flex items-center justify-between whitespace-nowrap border-b border-solid border-b-[#f0f2f4] px-10 py-3">
      <div className="flex items-center gap-4 text-[#111418]">
        <div className="size-4">
          <svg
            viewBox="0 0 48 48"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fillRule="evenodd"
              clipRule="evenodd"
              d="M24 4H42V17.3333V30.6667H24V44H6V30.6667V17.3333H24V4Z"
              fill="currentColor"
            ></path>
          </svg>
        </div>
        <h2 className="text-[#111418] text-lg font-bold leading-tight tracking-[-0.015em]">
          AgriSupport
        </h2>
      </div>
      <div className="flex flex-1 justify-end gap-8">
        <div className="flex items-center gap-9">
          <Link
            className="text-[#111418] text-sm font-medium leading-normal"
            to="/"
          >
            ホーム
          </Link>
          <Link
            className="text-[#111418] text-sm font-medium leading-normal"
            to="/disasters"
          >
            災害情報
          </Link>
          <Link
            className="text-[#111418] text-sm font-medium leading-normal"
            to="/application"
          >
            申請
          </Link>
          <Link
            className="text-[#111418] text-sm font-medium leading-normal"
            to="/facility-equipment"
          >
            施設設備
          </Link>
          <Link
            className="text-[#111418] text-sm font-medium leading-normal"
            to="/damage-levels"
          >
            被害程度
          </Link>
          <Link
            className="text-[#111418] text-sm font-medium leading-normal"
            to="/users"
          >
            ユーザー
          </Link>
          <Link
            className="text-[#111418] text-sm font-medium leading-normal"
            to="/organizations"
          >
            組織
          </Link>
        </div>
        <button className="flex max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 bg-[#f0f2f4] text-[#111418] gap-2 text-sm font-bold leading-normal tracking-[0.015em] min-w-0 px-2.5">
          <div className="text-[#111418]">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="20px"
              height="20px"
              fill="currentColor"
              viewBox="0 0 256 256"
            >
              <path d="M221.8,175.94C216.25,166.38,208,139.33,208,104a80,80,0,1,0-160,0c0,35.34-8.26,62.38-13.81,71.94A16,16,0,0,0,48,200H88.81a40,40,0,0,0,78.38,0H208a16,16,0,0,0,13.8-24.06ZM128,216a24,24,0,0,1-22.62-16h45.24A24,24,0,0,1,128,216ZM48,184c7.7-13.24,16-43.92,16-80a64,64,0,1,1,128,0c0,36.05,8.28,66.73,16,80Z"></path>
            </svg>
          </div>
        </button>
        <div
          className="bg-center bg-no-repeat aspect-square bg-cover rounded-full size-10"
          style={{
            backgroundImage:
              'url("https://lh3.googleusercontent.com/aida-public/AB6AXuCm5IN8QKkL-lJx0Q9_mG2_4r9vIEtptFifAGkNomjcCGwiEftXFxA_n05l6T8XjBGLHSk8PtOM83D9DnfJ-R9OLKjOsOyU5fRCf_Ef5_vPhHitQuj5Q0yubOWYjv1yhQ3vNi051H63y2nsOojqs93D_uHp9zXAyHeiD68KhIUyqtqISYc5w0IAOjrlT8NxHHYSrWVA5mDvEnDfLtYqcZCknU662Ur2WUub3LUsYg47yLnrPfzTY5HqeN3e2skAaKdnGKHsLhOcy70")',
          }}
        ></div>
      </div>
    </header>
  )
}

export const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          {children}
        </div>
      </main>
    </div>
  )
}
