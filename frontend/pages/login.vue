<script setup lang="ts"></script>
<template>
  <div class="antialiased">
    <div
      class="max-w-lg mx-auto my-10 bg-white p-8 rounded-xl shadow shadow-slate-400"
    >
      <h1 class="text-4xl font-medium">Login</h1>
      <p class="text-slate-500">Hi, Welcome back 👋</p>

      <div class="my-5">
        <button
          class="w-full text-center py-3 my-3 border flex space-x-2 items-center justify-center border-slate-200 rounded-lg text-slate-700 hover:border-slate-400 hover:text-slate-900 hover:shadow transition duration-150"
        >
          <img
            src="https://www.svgrepo.com/show/355037/google.svg"
            class="w-6 h-6"
            alt=""
          />
          <span>Login with Google</span>
        </button>
      </div>
      <form action="" class="my-10">
        <div class="flex flex-col space-y-5">
          <label for="email">
            <p class="font-medium text-slate-700 pb-2">Email</p>
            <input
              id="email"
              name="email"
              type="email"
              class="w-full py-3 border border-slate-200 rounded-lg px-3 focus:outline-none focus:border-slate-500 hover:shadow"
              placeholder="メールアドレスを入力してください"
            />
          </label>
          <label for="password">
            <p class="font-medium text-slate-700 pb-2">Password</p>
            <input
              id="password"
              name="password"
              type="password"
              class="w-full py-3 border border-slate-200 rounded-lg px-3 focus:outline-none focus:border-slate-500 hover:shadow"
              placeholder="パスワードを入力してください"
            />
          </label>

          <button
            class="w-full py-3 font-medium text-white bg-indigo-600 hover:bg-indigo-500 rounded-lg border-indigo-500 hover:shadow inline-flex space-x-2 items-center justify-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"
              />
            </svg>
            <span>Login</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style>
:root {
  font-family: "Inter", sans-serif;
}
@supports (font-variation-settings: normal) {
  :root {
    font-family: "Inter var", sans-serif;
  }
}
</style>
