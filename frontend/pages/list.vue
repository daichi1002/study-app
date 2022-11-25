<template>
  <div class="tw-container tw-mx-auto">
    <div class="tw-flex tw-justify-center tw-mt-10">
      <input
        class="tw-bg-gray-200 tw-appearance-none tw-border-2 tw-border-gray-200 tw-rounded tw-w-2/4 tw-py-2 tw-px-4 tw-text-gray-700 tw-leading-tight focus:tw-outline-none focus:tw-bg-white focus:tw-border-purple-500"
        id="inline-full-name"
        type="text"
        value=""
      />
      <button>
        <img src="../assets/images/search.png" />
      </button>
    </div>
    <div class="tw-flex tw-justify-center tw-mt-10">
      <table class="tw-table-auto">
        <thead>
          <tr>
            <th class="tw-px-4 tw-py-2">タイトル</th>
            <th class="tw-px-4 tw-py-2">タグ</th>
            <th class="tw-px-4 tw-py-2">投稿者</th>
            <th class="tw-px-4 tw-py-2">更新日</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td class="tw-border tw-px-4 tw-py-2">Intro to CSS</td>
            <td class="tw-border tw-px-4 tw-py-2">Adam</td>
            <td class="tw-border tw-px-4 tw-py-2">858</td>
            <td class="tw-border tw-px-4 tw-py-2">2022/08/10</td>
          </tr>
          <tr class="tw-bg-gray-200">
            <td class="tw-border tw-px-4 tw-py-2">
              A Long and Winding Tour of the History of UI Frameworks and Tools
              and the Impact on Design
            </td>
            <td class="tw-border tw-px-4 tw-py-2">Adam</td>
            <td class="tw-border tw-px-4 tw-py-2">112</td>
            <td class="tw-border tw-px-4 tw-py-2">2022/08/10</td>
          </tr>
          <tr>
            <td class="tw-border tw-px-4 tw-py-2">Intro to JavaScript</td>
            <td class="tw-border tw-px-4 tw-py-2">Chris</td>
            <td class="tw-border tw-px-4 tw-py-2">1,280</td>
            <td class="tw-border tw-px-4 tw-py-2">2022/08/10</td>
          </tr>
        </tbody>
      </table>
    </div>
    <button
      class="tw-fixed tw-bottom-8 tw-right-8 tw-rounded-full tw-w-24 tw-py-6 tw-bg-teal-300 hover:tw-bg-teal-500 tw-text-white tw-text-center tw-font-semibold tw-text-5xl"
    >
      +
    </button>
  </div>
</template>
