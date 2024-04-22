import React, { useState, useEffect } from 'react';
import axios from 'axios';

const AllPostsrender = () => {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/posts');
        setPosts(response.data);
      } catch (error) {
        console.error('Error fetching posts:', error);
        // Handle errors appropriately (e.g., display error message to user)
      }
    };

    fetchPosts();
  }, []); // Run effect only once on component mount

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