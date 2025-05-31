import { useUsers } from "../hooks/useUsers"
import { UserList } from "./UserList"

export const UserFeature = () => {
  // ユーザー情報を取得
  const { data: users, isLoading, error } = useUsers()

  return (
    <>
      <div className="flex flex-wrap justify-between gap-3 p-4">
        <div className="flex min-w-72 flex-col gap-3">
          <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            ユーザー一覧
          </p>
          <p className="text-[#637588] text-sm font-normal leading-normal">
            システムに登録されているユーザーの一覧です。
          </p>
        </div>
      </div>

      {/* ユーザーリスト */}
      <UserList 
        users={users} 
        isLoading={isLoading} 
        error={error} 
      />
    </>
  )
}