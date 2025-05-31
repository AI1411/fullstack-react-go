"use client"

import { useListDisasters } from "@/api/client"
import Header from "@/components/layout/header/page"
import Link from "next/link"
import { useState } from "react"
import SupportApplicationModal from "@/components/application/SupportApplicationModal"

export default function Home() {
  const {
    data: disastersResponse,
    isLoading,
    error,
  } = useListDisasters({
    query: {
      staleTime: 5 * 60 * 1000, // 5 minutes cache
    },
  })

  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleOpenModal = () => setIsModalOpen(true);
  const handleCloseModal = () => setIsModalOpen(false);
  const handleApplicationCreated = () => {
    console.log("New support application created successfully from Home page!");
    alert("Application created successfully."); // Changed to English
    // Example: queryClient.invalidateQueries('supportApplications');
  };

  if (isLoading) {
    return <div>Loading...</div> // Changed to English
  }

  if (error) {
    return <div>Failed to fetch user data.</div> // Changed to English and assuming it's user data problem
  }

  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <div className="flex flex-wrap justify-between gap-3 p-4 items-center">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                Dashboard
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                Welcome! Check the system overview and latest disaster information.
              </p>
            </div>
            <div className="flex-shrink-0">
              <button
                onClick={handleOpenModal}
                className="inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-colors"
              >
                Register Support Application
              </button>
            </div>
          </div>

          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            System Overview
          </h2>
          <div className="p-4">
            <p className="text-gray-700">Information about the system overview will be displayed here.</p>
          </div>

          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            Latest Disaster Information
          </h2>
          <div className="p-4">
            {disastersResponse && disastersResponse.disasters && disastersResponse.disasters.length > 0 ? (
              <ul className="space-y-2">
                {disastersResponse.disasters.slice(0, 3).map((disaster) => (
                  <li key={disaster.id} className="p-3 bg-gray-100 rounded-md shadow-sm">
                    <h3 className="font-semibold">{disaster.disaster_name}</h3>
                    <p className="text-sm text-gray-600">{disaster.start_date}</p>
                  </li>
                ))}
              </ul>
            ) : (
              <p className="text-gray-700">No disaster information currently reported.</p>
            )}
          </div>

          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            Recent Disasters List
          </h2>
          {/* Content for recent disasters list */}

          <div className="p-4 text-center">
            <Link
              href="/disasters"
              className="inline-flex items-center justify-center px-6 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-[#111418] hover:bg-[#333] transition-colors"
            >
              View All Disaster Information
            </Link>
          </div>
        </div>
      </main>
      <SupportApplicationModal
        isOpen={isModalOpen}
        onClose={handleCloseModal}
        onApplicationCreated={handleApplicationCreated}
      />
    </div>
  )
}
