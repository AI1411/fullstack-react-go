import { Link, useParams } from "@tanstack/react-router"
import { useUser } from "../hooks/useUser"
import { formatDate } from "../../../utils/formatters"

export const UserDetail = () => {
  // Get the user ID from the URL parameters
  const { userId } = useParams({ from: "/users/$userId" })

  // Fetch the user details
  const { data: user, isLoading, error } = useUser(userId)

  if (isLoading) {
    return <div className="p-4">読み込み中...</div>
  }

  if (error) {
    return <div className="p-4">ユーザー情報の取得に失敗しました</div>
  }

  if (!user) {
    return <div className="p-4">ユーザーが見つかりませんでした</div>
  }

  return (
    <div className="p-4">
      <div className="mb-4">
        <Link
          to="/users"
          className="text-[#197fe5] hover:underline flex items-center"
        >
          ← ユーザーリストに戻る
        </Link>
      </div>

      <div className="flex flex-wrap justify-between gap-3 mb-6">
        <div className="flex min-w-72 flex-col gap-3">
          <h1 className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            {user.name}
          </h1>
        </div>
      </div>

      <div className="bg-white rounded-lg border border-[#dce0e5] overflow-hidden mb-6">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-6">
          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">ID</h3>
            <p className="text-[#111418] text-base font-normal">
              {user.id}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">名前</h3>
            <p className="text-[#111418] text-base font-normal">
              {user.name}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">メールアドレス</h3>
            <p className="text-[#111418] text-base font-normal">
              {user.email}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">作成日</h3>
            <p className="text-[#111418] text-base font-normal">
              {user.created_at ? formatDate(user.created_at) : "不明"}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">更新日</h3>
            <p className="text-[#111418] text-base font-normal">
              {user.updated_at ? formatDate(user.updated_at) : "不明"}
            </p>
          </div>
        </div>
      </div>

      {user.organizations && user.organizations.length > 0 && (
        <div className="bg-white rounded-lg border border-[#dce0e5] overflow-hidden mb-6">
          <div className="p-6">
            <h2 className="text-[#111418] text-xl font-semibold mb-4">所属組織</h2>
            <div className="grid grid-cols-1 gap-4">
              {user.organizations.map((org) => (
                <div key={org.id} className="flex items-center p-3 border border-[#dce0e5] rounded-md">
                  <div className="flex-1">
                    <p className="text-[#111418] text-base font-medium">{org.name}</p>
                    <p className="text-[#637588] text-sm">{org.type}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  )
}
