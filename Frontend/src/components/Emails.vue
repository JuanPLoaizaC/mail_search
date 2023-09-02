<template>
  <div class="grid grid-flow-row auto-rows-max pr-8 pl-8" style="margin-top: 2em">
    <label class="relative block">
      <span class="sr-only">Search</span>
      <span class="absolute inset-y-0 left-0 flex items-center pl-2">
        <svg class="h-5 w-5 fill-slate-300" viewBox="0 0 20 20"><!-- ... --></svg>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
          <path strokeLinecap="round" strokeLinejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
        </svg>
      </span>
      <input class="placeholder:italic placeholder:text-slate-400 block bg-white w-full border border-slate-300 rounded-md py-2 pl-9 pr-3 shadow-sm focus:outline-none focus:border-sky-500 focus:ring-sky-500 focus:ring-1 sm:text-sm font-sans md:italic" placeholder="Search for anything..." type="text" name="search"/>
    </label>
  </div>
  <div class="flex flex-col" style="margin-top: 1em;">
    <div class="overflow-x-auto">
      <div class="p-1.5 w-full inline-block align-middle">
        <div class="overflow-hidden border rounded-lg">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase"
                >
                  Subject
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase"
                >
                  From
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase"
                >
                  To
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr class="hover:bg-slate-100">
                <td
                  class="px-6 py-4 text-sm font-medium text-gray-800 whitespace-nowrap"
                >
                  1
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  Jone Doe
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  jonne62@gmail.com
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>