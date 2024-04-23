import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

function AllPosts() {
    const [posts, setPosts] = useState([]);
    const [errorMessage, setErrorMessage] = useState('');

    useEffect(() => {
        fetchAllPosts();
    }, []);

    const fetchAllPosts = async () => {
        try {
            const response = await axios.get(
                'http://localhost:8080/api/posts/all'
            );
            setPosts(response.data);
        } catch (error) {
            console.error('Error fetching all posts:', error);
            setErrorMessage('Error fetching all posts. Please try again later.');
        }
    };

    return (
        <div className="max-w-md mx-auto my-8 p-6 bg-white rounded-lg shadow-md">
            <h2 className="text-2xl mb-4">All Posts</h2>
            {errorMessage && <p className="text-red-500 mb-4">{errorMessage}</p>}
            <ul>
                {posts.map((post) => (
                    <li key={post._id} className="mt-2">
                        <Link to={`/posts/${post._id}`} className="text-blue-500 hover:underline">{post.title}</Link>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default AllPosts;
