import React, { useState } from 'react';
import axios from 'axios'; // Assuming you've set up Axios for HTTP requests

const CreatePostForm = () => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Prepare data to send
    const data = {
      title,
      content,
    };

    try {
      // Send POST request to create a post
      const response = await axios.post('http://localhost:8080/api/posts', data);

      // Handle successful creation
      console.log('Post created successfully:', response.data);
      setTitle(''); // Clear form fields after successful creation (optional)
      setContent('');
    } catch (error) {
      console.error('Error creating post:', error);
      // Handle errors appropriately (e.g., display error message to user)
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col space-y-2">
      <div className=" mt-16 flex justify-center">
      <label htmlFor="title" className="text-sm font-medium">
        Title:
      </label>
      <input
        type="text"
        id="title"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        required
        className=" rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500"
      />
      </div>
      <div className="flex justify-center">
      <label htmlFor="content" className="text-sm font-medium">
        Content:
      </label>
      <textarea
        id="content"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        required
        className="w-1/2 rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500 h-24"
      />
      </div>

      <div className="flex justify-center">
      <button type="submit" className="items-center bg-blue-500 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-700">
        Create Post
      </button>
      </div>
      <div className="flex justify-center">
      <a
              href="/allposts"
              class="inline-flex justify-center items-center py-3 px-5 text-base font-medium text-center text-white rounded-lg bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-900"
            >
              View all posts
            </a>
            </div>
    </form>
  );
};

export default CreatePostForm;