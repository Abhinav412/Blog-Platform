const Home = () => {
    const handleSubmit = (event) => {
        event.preventDefault();
        const comment = event.target.comment.value;
        console.log(comment);
        event.target.reset();
    };

    return (
        <div>
            <nav className="bg-white dark:bg-gray-900 fixed w-full z-20 top-0 start-0 border-b border-gray-200 dark:border-gray-600">
                <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                    <a href="http://localhost:3000/Home" className="flex items-center space-x-3 rtl:space-x-reverse">
                        <img src="https://1000logos.net/wp-content/uploads/2017/01/DC-Comics-Logo.jpg" className="h-8" alt="comic"/>
                        <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">The DCU Blog</span>
                    </a>
                    {/* <div className="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
                        <button type="button" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Get started</button>
                        <button data-collapse-toggle="navbar-sticky" type="button" className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600" aria-controls="navbar-sticky" aria-expanded="false">
                            <span className="sr-only">Open main menu</span>
                            <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
                                <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M1 1h15M1 7h15M1 13h15" />
                            </svg>
                        </button>
                    </div> */}
                    {/* <div className="items-center justify-between hidden w-full md:flex md:w-auto md:order-1" id="navbar-sticky">
                        <ul className="py-8 px-4 mx-auto flex flex-col p-4 md:p-0 mt-4 font-medium border border-gray-100 rounded-lg bg-gray-50 md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                            <li>
                                <a href="#" className="block py-2 px-3 text-white bg-blue-700 rounded md:bg-transparent md:text-blue-700 md:p-0 md:dark:text-blue-500" aria-current="page">Home</a>
                            </li>
                            <li>
                                <a href="#" className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:p-0 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">About</a>
                            </li>
                            <li>
                                <a href="#" className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:p-0 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">Services</a>
                            </li>
                            <li>
                                <a href="#" className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:p-0 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">Contact</a>
                            </li>
                        </ul>
                    </div> */}
                </div>
            </nav>
            <div className="bg-indigo-50">
                <p className=" ml-16 w-1/2 text-left">
                    <br/>
                    <br/>
                    <div className=" text-4xl flex mt-16 mb-8 justify-center"><h1>SUPERMAN (11-07-2025)</h1></div>
                    <div className="text-lg">
                    "It focuses on Superman balancing his Kryptonian heritage with his human upbringing," Safran said. "He is the embodiment of truth, justice and the American way. He is kindness in a world that thinks of kindness as old-fashioned."
                    <br/>
                    <br/>
                    "I really love the idea of Superman. He's a big old galoot," Gunn said. "He is a farm boy from Kansas who is very idealistic. His greatest weakness is that he'll never kill anybody, doesn't want to hurt a living soul. And I like that sort of innate goodness about Superman as his defining characteristic." 
                    <br/>
                    </div>

                </p>

                <form onSubmit={handleSubmit}>
                    <div className=" py-12 px-4 mb-4 rounded-lg dark:bg-gray-700 dark:border-gray-600">
                        <div className="w-1/2 px-4 py-2 ml-12 rounded-t-lg dark:bg-gray-800">
                            <label htmlFor="comment" className="sr-only">Your comment</label>
                            <textarea id="comment" name="comment" rows="4" className="bg-white h-11 w-full px-0 text-sm text-gray-900 border-0 dark:bg-gray-800 focus:ring-0 dark:text-white dark:placeholder-gray-400" placeholder="Write a comment..." required></textarea>  
                        </div>

                        <div className=" ml-12 px-3 py-2 dark:border-gray-600">
                            <button type="submit" className="inline-flex items-center py-2.5 px-4 text-xs font-medium text-center text-white bg-slate-600 rounded-lg focus:ring-4 focus:ring-blue-200 dark:focus:ring-blue-900 hover:bg-blue-800">
                                Post comment
                            </button>
                        </div>
                    </div>
                </form>
            </div>
            <section class="h-full w-full">
            <div class="bg-red-100 flex flex-col items-center justify-center px-6 py-16 mx-auto md:h-screen lg:py-0">
                <div class="w-full bg-white rounded-lg shadow dark:border dark:border-gray-700 md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800">
                    <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                    </div>
                </div>
            </div>
            </section>   
        </div>
    );
};

export default Home;
