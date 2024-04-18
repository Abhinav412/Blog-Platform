import React, { useState, useEffect } from 'react';
import axios from 'axios';

const AllPosts = () => {
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
    <div className="container mx-auto max-w-lg">
      <h1 className="text-2xl font-bold mb-4">All Posts</h1>
      {posts.length > 0 ? (
        <ul>
          {posts.map((post) => (
            <li key={post.ID} className="border border-gray-300 p-3 mb-2">
              <h3>{post.title}</h3>
              <p>{post.content}</p>
            </li>
          ))}
        </ul>
      ) : (
        <p>No posts found.</p>
      )}
    </div>
  );
};

export default AllPostsrender;