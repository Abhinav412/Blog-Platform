const Welcome = () =>  {
    return (
      <section>
        <div className="bg-cover bg-center bg-no-repeat h-screen w-screen flex flex-col items-center justify-center" style={{ backgroundImage: "url('https://images.unsplash.com/photo-1696185570507-2d1283399560?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D')" }}>
        <div class="py-8 px-4 mx-auto max-w-screen-xl text-center lg:py-16">
          <h1 class="mb-4 text-4xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-6xl dark:text-white">
            Welcome to the Blog
          </h1>
          <br/>
          <p class="mb-8 font-normal text-gray-500 lg:text-3xl sm:px-16 lg:px-48 dark:text-gray-400">
            This is Go Lang Project
          </p>
          <br/>
          <div class="flex flex-col space-y-4 sm:flex-row sm:justify-center sm:space-y-0 sm:space-x-4">
            <a
              href="/createpost"
              class="inline-flex justify-center items-center py-3 px-5 text-base font-medium text-center text-white rounded-lg bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-900"
            >
              Get Started
            </a>
          </div>
        </div>
        </div>
      </section>
    );
}
  
export default Welcome;
  