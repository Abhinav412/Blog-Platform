import React, { useState, useEffect } from 'react';
import axios from 'axios';

function CreatePost() {
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    const [author, setAuthor] = useState('');
    const [lastPostAuthor, setLastPostAuthor] = useState('');
    const [userPosts, setUserPosts] = useState([]);
    const [errorMessage, setErrorMessage] = useState('');

    useEffect(() => {
        if (lastPostAuthor) {
            fetchUserPosts(lastPostAuthor);
        }
    }, [lastPostAuthor]);

    const fetchUserPosts = async (authorName) => {
        try {
            const response = await axios.get(
                `http://localhost:8080/api/posts?author=${authorName}`
            );
            setUserPosts(response.data);
        } catch (error) {
            console.error('Error fetching user posts:', error);
            setErrorMessage('Error fetching user posts. Please try again later.');
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post(
                'http://localhost:8080/api/posts',
                {
                    title: title,
                    content: content,
                    author: author,
                }
            );
            console.log('Post created:', response.data);
            setLastPostAuthor(author); // Update the author of the last created post
            // Clear form fields after successful submission
            setTitle('');
            setContent('');
            setAuthor('');
            setErrorMessage('');
        } catch (error) {
            console.error('Error creating post:', error);
            setErrorMessage('Error creating post. Please try again later.');
        }
    };

    const handleViewUserPosts = async () => {
        try {
            const response = await axios.get(
                `http://localhost:8080/api/posts?author=${lastPostAuthor}`
            );
            setUserPosts(response.data);
        } catch (error) {
            console.error('Error fetching user posts:', error);
            setErrorMessage('Error fetching user posts. Please try again later.');
        }
    };

    const handleDeletePost = async (postId) => {
        try {
            await axios.delete(`http://localhost:8080/api/posts/${postId}`);
            setUserPosts(userPosts.filter((post) => post._id !== postId));
        } catch (error) {
            console.error('Error deleting post:', error);
            setErrorMessage('Error deleting post. Please try again later.');
        }
    };

    return (
        <div className="max-w-md mx-auto my-8 p-6 bg-white rounded-lg shadow-md">
            <h2 className="text-2xl mb-4">Create a New Post</h2>
            {errorMessage && <p className="text-red-500 mb-4">{errorMessage}</p>}
            <form onSubmit={handleSubmit}>
                <div className="mb-4">
                    <label className="block text-gray-700 mb-2">Title:</label>
                    <input
                        type="text"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        required
                        className="block w-full border border-gray-300 rounded-md p-2"
                    />
                </div>
                <div className="mb-4">
                    <label className="block text-gray-700 mb-2">Content:</label>
                    <textarea
                        value={content}
                        onChange={(e) => setContent(e.target.value)}
                        required
                        className="block w-full border border-gray-300 rounded-md p-2"
                    ></textarea>
                </div>
                <div className="mb-4">
                    <label className="block text-gray-700 mb-2">Author:</label>
                    <input
                        type="text"
                        value={author}
                        onChange={(e) => setAuthor(e.target.value)}
                        required
                        className="block w-full border border-gray-300 rounded-md p-2"
                    />
                </div>
                <button type="submit" className="bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 transition duration-300">Create Post</button>
            </form>
            {lastPostAuthor && (
                <button onClick={handleViewUserPosts} className="mt-4 bg-gray-200 text-gray-800 py-2 px-4 rounded-md hover:bg-gray-300 transition duration-300">View User's Posts</button>
            )}
            <h2 className="text-xl mt-8">User's Posts</h2>
            <ul>
                {userPosts.map((post) => (
                    <li key={post._id} className="mt-2">
                        <strong>{post.title}</strong>: {post.content}
                        <button onClick={() => handleDeletePost(post._id)} className="ml-2 bg-red-500 text-white py-1 px-2 rounded-md hover:bg-red-600 transition duration-300">Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default CreatePost;


