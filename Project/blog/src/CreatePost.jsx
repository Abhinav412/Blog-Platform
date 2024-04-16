// CreatePost.js
import React, { useState } from 'react';

const CreatePost = () => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('/api/posts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title, content }),
      });
      if (response.ok) {
        // Post created successfully
        console.log('Post created successfully');
        // Redirect user to a different page, or show a success message
      } else {
        // Handle error
        console.error('Error creating post');
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <section>
      <div className="bg-gray-100 min-h-screen flex flex-col items-center justify-center">
        <div className="py-8 px-4 mx-auto max-w-screen-xl text-center">
          <h1 className="mb-4 text-4xl font-extrabold text-gray-900">
            Create a New Blog Post
          </h1>
          <form onSubmit={handleSubmit} className="max-w-md mx-auto">
            <div className="mb-4">
              <label htmlFor="title" className="block text-lg text-gray-700">
                Title
              </label>
              <input
                type="text"
                id="title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                className="mt-1 px-4 py-2 w-full border rounded-md"
                required
              />
            </div>
            <div className="mb-4">
              <label htmlFor="content" className="block text-lg text-gray-700">
                Content
              </label>
              <textarea
                id="content"
                value={content}
                onChange={(e) => setContent(e.target.value)}
                className="mt-1 px-4 py-2 w-full border rounded-md"
                rows="6"
                required
              ></textarea>
            </div>
            <button
              type="submit"
              className="inline-block py-3 px-6 text-lg font-semibold text-white bg-blue-700 rounded-md hover:bg-blue-800"
            >
              Create Post
            </button>
            <a
              href="/allposts"
              class="inline-flex justify-center items-center py-3 px-5 text-base font-medium text-center text-white rounded-lg bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-900"
            >
              View all posts 
            </a>
           
          </form>
        </div>
      </div>
    </section>
  );
};

export default CreatePost;
