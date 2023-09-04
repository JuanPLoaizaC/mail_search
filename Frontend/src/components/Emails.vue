<template>
  <div class="flex flex-col pr-8 pl-8" style="margin-top: 1em;">
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