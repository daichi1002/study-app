<template>
  <div class="container mx-auto">
    <div class="flex justify-center mt-10">
      <input
        class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-2/4 py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
        id="inline-full-name"
        type="text"
        value=""
      />
      <button>
        <img src="../assets/images/search.png" />
      </button>
    </div>
    <div class="flex justify-center mt-10">
      <table class="table-auto">
        <thead>
          <tr>
            <th class="px-4 py-2">タイトル</th>
            <th class="px-4 py-2">タグ</th>
            <th class="px-4 py-2">投稿者</th>
            <th class="px-4 py-2">更新日</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td class="border px-4 py-2">Intro to CSS</td>
            <td class="border px-4 py-2">Adam</td>
            <td class="border px-4 py-2">858</td>
            <td class="border px-4 py-2">2022/08/10</td>
          </tr>
          <tr class="bg-gray-200">
            <td class="border px-4 py-2">
              A Long and Winding Tour of the History of UI Frameworks and Tools
              and the Impact on Design
            </td>
            <td class="border px-4 py-2">Adam</td>
            <td class="border px-4 py-2">112</td>
            <td class="border px-4 py-2">2022/08/10</td>
          </tr>
          <tr>
            <td class="border px-4 py-2">Intro to JavaScript</td>
            <td class="border px-4 py-2">Chris</td>
            <td class="border px-4 py-2">1,280</td>
            <td class="border px-4 py-2">2022/08/10</td>
          </tr>
        </tbody>
      </table>
    </div>
    <button
      class="fixed bottom-8 right-8 rounded-full w-24 py-6 bg-teal-300 hover:bg-teal-500 text-white text-center font-semibold text-5xl"
    >
      +
    </button>
  </div>
</template>
