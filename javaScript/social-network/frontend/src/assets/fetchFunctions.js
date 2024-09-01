export async function getUser(userId) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/user?id=${userId}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get user");
    }

    const data = await response.json();
    return data.user;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getUsers() {
  try {
    const response = await fetch("http://localhost:3000/api/v1/users", {
      method: "GET",
      credentials: "include",
    });

    if (!response.ok) {
      throw new Error("Failed to get users");
    }

    const data = await response.json();
    return data.users;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function updateUser(privacy) {
  try {
    const response = await fetch(`http://localhost:3000/api/v1/user/update`, {
      method: "PUT",
      credentials: "include",
      body: JSON.stringify({
        privacy: privacy,
      }),
    });

    if (!response.ok) {
      throw new Error("Failed to update user");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function followUser(userId, follow) {
  try {
    const response = await fetch(`http://localhost:3000/api/v1/user/follow`, {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        userId: userId,
        follow: follow,
      }),
    });

    if (!response.ok) {
      throw new Error("Failed to update user");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getPosts(offset) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/posts?offset=${offset}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get posts");
    }

    const data = await response.json();
    return data.posts;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getUserPosts(userId) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/posts/user?id=${userId}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get posts");
    }

    const data = await response.json();
    return data.posts;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getUserFollowers(userId) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/user/followers?id=${userId}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get posts");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getNotifications() {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/user/notifications`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get posts");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function updateFollow(userId, decision) {
  try {
    const response = await fetch(`http://localhost:3000/api/v1/follow/update`, {
      method: "PUT",
      credentials: "include",
      body: JSON.stringify({
        userId: userId,
        decision: decision,
      }),
    });

    if (!response.ok) {
      throw new Error("Failed to get posts");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function sendMessage(toId, message, user) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/message/create`,
      {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({
          toId: toId,
          fromId: user ? 1 : 0,
          message: message,
        }),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to send message");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getLastMessage(userId) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/message?id=${userId}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get message");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getLastGroupMessage(groupId) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/group/message?id=${groupId}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get message");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getMessages(userId, offset) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/messages?id=${userId}&offset=${offset}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get messages");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getGroupMessages(groupId, offset) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/group/messages?id=${groupId}&offset=${offset}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get messages");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getGroup(id) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/group?id=${id}`,
      {
        method: "GET",
        credentials: "include",
      }
    );

    if (!response.ok) {
      throw new Error("Failed to get groups");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getGroups() {
  try {
    const response = await fetch(`http://localhost:3000/api/v1/groups`, {
      method: "GET",
      credentials: "include",
    });

    if (!response.ok) {
      throw new Error("Failed to get groups");
    }

    const data = await response.json();
    return data.groups;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function createEvent(data) {
  try {
    const response = await fetch(`http://localhost:3000/api/v1/event/create`, {
      method: "POST",
      credentials: "include",
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error("Failed to create event");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function attendEvent(data) {
  try {
    const response = await fetch(`http://localhost:3000/api/v1/event/going`, {
      method: "POST",
      credentials: "include",
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error("Failed to update attendance info");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function createGroupUser(data) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/group/user/create`,
      {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(data),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to create group user");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function updateGroupUser(data) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/v1/group/user/update`,
      {
        method: "PUT",
        credentials: "include",
        body: JSON.stringify(data),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to update group user");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}
