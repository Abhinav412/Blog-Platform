// AllPosts.js
import React, { useState, useEffect } from 'react';

const AllPosts = () => {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    fetch('/api/allposts')
      .then(response => response.json())
      .then(data => setPosts(data))
      .catch(error => console.error('Error fetching posts:', error));
  }, []);

  return (
    <section>
          <nav className="bg-white dark:bg-gray-900 fixed w-full z-20 top-0 start-0 border-b border-gray-200 dark:border-gray-600">
            <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
              <a href="http://localhost:3000/Home" className="flex items-center space-x-3 rtl:space-x-reverse">
                  <img src="https://1000logos.net/wp-content/uploads/2017/01/DC-Comics-Logo.jpg" className="h-8" alt="comic"/>
                  <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">The DCU Blog</span>
              </a>
              </div>
          </nav>

      <div className="bg-gray-100 min-h-screen flex flex-col items-center justify-center">
        <div className="py-8 px-4 mx-auto max-w-screen-xl text-center">
          <h1 className="mb-4 text-4xl font-extrabold text-gray-900">
            All Blog Posts
          </h1>
          <div className="max-w-md mx-auto">
            {posts.map(post => (
              <div key={post.id} className="mb-8 border rounded-lg p-4">
                <h2 className="text-2xl font-bold mb-2">{post.title}</h2>
                <p className="text-gray-700">{post.content}</p>
              </div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};

export default AllPosts;
