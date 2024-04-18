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
                    <div className=" py-12 px-4 rounded-lg dark:bg-gray-700 dark:border-gray-600">
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
        </div>
    );
};

export default Home;
