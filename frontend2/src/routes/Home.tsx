import { useListDisasters } from "../api/generated/client"
import { Link } from "@tanstack/react-router"

export const Home = () => {
  // 災害情報を取得
  const {
    data: disastersResponse,
    isLoading,
    error,
  } = useListDisasters({
    query: {
      staleTime: 5 * 60 * 1000, // 5分間キャッシュ
    },
  })

  if (isLoading) {
    return <div>読み込み中</div>
  }

  if (error) {
    return <div>ユーザーの取得に失敗しました</div>
  }

  return (
    <>
      <div className="flex flex-wrap justify-between gap-3 p-4">
        <div className="flex min-w-72 flex-col gap-3">
          <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            ダッシュボード
          </p>
          <p className="text-[#637588] text-sm font-normal leading-normal">
            お疲れさまです！システムの概要と最新の災害情報をご確認ください。
          </p>
        </div>
      </div>

      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        システム概要
      </h2>
      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        最新災害情報
      </h2>

      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        最近の災害一覧
      </h2>

      {/* 全件表示リンク */}
      <div className="p-4 text-center">
        <Link
          to="/disasters"
          className="inline-flex items-center justify-center px-6 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-[#111418] hover:bg-[#333] transition-colors mr-4"
        >
          すべての災害情報を見る
        </Link>
        <Link
          to="/application"
          className="inline-flex items-center justify-center px-6 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-[#197fe5] hover:bg-[#1565c0] transition-colors"
        >
          支援申請を行う
        </Link>
      </div>
    </>
  )
}
