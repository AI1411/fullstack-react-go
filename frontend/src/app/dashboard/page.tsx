import Header from "@/components/layout/header/page"

export default function Home() {
  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <div className="flex flex-wrap justify-between gap-3 p-4">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                Dashboard
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                Welcome back, Ms. Tanaka! Here's an overview of your tasks and
                system status.
              </p>
            </div>
          </div>
          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            System Overview
          </h2>
          <div className="flex flex-wrap gap-4 p-4">
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                Total Applications
              </p>
              <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                1,250
              </p>
            </div>
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                Pending Approvals
              </p>
              <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                320
              </p>
            </div>
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                Active Users
              </p>
              <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                150
              </p>
            </div>
          </div>
          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            Latest Disaster Information
          </h2>
          <div className="p-4">
            <div className="flex items-stretch justify-between gap-4 rounded-lg">
              <div className="flex flex-col gap-1 flex-[2_2_0px]">
                <p className="text-[#111418] text-base font-bold leading-tight">
                  Typhoon No. 10 Impact Assessment
                </p>
                <p className="text-[#637588] text-sm font-normal leading-normal">
                  Detailed report on the agricultural damage caused by Typhoon
                  No. 10, including affected areas and estimated recovery costs.
                </p>
              </div>
              <div
                className="w-full bg-center bg-no-repeat aspect-video bg-cover rounded-lg flex-1"
                style={{
                  backgroundImage:
                    'url("https://lh3.googleusercontent.com/aida-public/AB6AXuAwL6_AdRbGqk3fz9oAyKgApsJ5lzCZr323vDQidQ9sUfYW8fL05o-F1utFzuhac0AdevlWakVlW9vzMCRB7o_50MQ7boxvgVAkfcpYppzmOj0ApvLQc-dIfIhILwFZEzbaAXyFtfO4opsZF3lJTppgRsbw5Bs-DkYdzhVUAOh0Azxj54F00OhUq-XNDvnuNK5PCypb7MbFNqq1njXjjA8Mze_JUSLaovZSz4hDcO_wGxaoLBjVYxHtVpnfSQdmZDYXQmn10XKZ7kY")',
                }}
              ></div>
            </div>
          </div>
          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            Application Status
          </h2>
          <div className="px-4 py-3 @container">
            <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
              <table className="flex-1">
                <thead>
                  <tr className="bg-white">
                    <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Application ID
                    </th>
                    <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Applicant Name
                    </th>
                    <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-60 text-sm font-medium leading-normal">
                      Damage Type
                    </th>
                    <th className="table-column-480 px-4 py-3 text-left text-[#111418] w-60 text-sm font-medium leading-normal">
                      Status
                    </th>
                    <th className="table-column-600 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Submission Date
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {/* Table rows data */}
                  <tr className="border-t border-t-[#dce0e5]">
                    <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                      APP2023-001
                    </td>
                    <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      Mr. Sato
                    </td>
                    <td className="table-column-360 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Crop Damage</span>
                      </button>
                    </td>
                    <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Pending</span>
                      </button>
                    </td>
                    <td className="table-column-600 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      2023-08-15
                    </td>
                  </tr>
                  {/* ... other rows ... */}
                  <tr className="border-t border-t-[#dce0e5]">
                    <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                      APP2023-002
                    </td>
                    <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      Ms. Suzuki
                    </td>
                    <td className="table-column-360 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Facility Damage</span>
                      </button>
                    </td>
                    <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Approved</span>
                      </button>
                    </td>
                    <td className="table-column-600 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      2023-08-16
                    </td>
                  </tr>
                  <tr className="border-t border-t-[#dce0e5]">
                    <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                      APP2023-003
                    </td>
                    <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      Mr. Takahashi
                    </td>
                    <td className="table-column-360 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Livestock Loss</span>
                      </button>
                    </td>
                    <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Rejected</span>
                      </button>
                    </td>
                    <td className="table-column-600 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      2023-08-17
                    </td>
                  </tr>
                  <tr className="border-t border-t-[#dce0e5]">
                    <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                      APP2023-004
                    </td>
                    <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      Ms. Ito
                    </td>
                    <td className="table-column-360 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Crop Damage</span>
                      </button>
                    </td>
                    <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Pending</span>
                      </button>
                    </td>
                    <td className="table-column-600 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      2023-08-18
                    </td>
                  </tr>
                  <tr className="border-t border-t-[#dce0e5]">
                    <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                      APP2023-005
                    </td>
                    <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      Mr. Watanabe
                    </td>
                    <td className="table-column-360 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Facility Damage</span>
                      </button>
                    </td>
                    <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                      <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                        <span className="truncate">Approved</span>
                      </button>
                    </td>
                    <td className="table-column-600 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                      2023-08-19
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}
