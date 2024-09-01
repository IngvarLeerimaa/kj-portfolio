INSERT INTO users (email, password, first_name, last_name, date_of_birth, avatar, nickname, about, public)
VALUES 
("user1@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "John", "Smith", "1990-03-15", "avatar.png", "johnny", "Tech enthusiast and movie buff.", 1),
("user2@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Lisa", "Johnson", "1988-07-20", "example1.jpg", NULL, "Coffee lover and bookworm.", 0),
("user3@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Michael", "Williams", "1992-11-30", "example6.jpg", "mikey", "Outdoor enthusiast and adventure seeker.", 1),
("user4@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Sarah", "Brown", "1985-04-05", "example3.jpg", "sarahb", "Foodie and travel enthusiast.", 0),
("user5@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "David", "Miller", "1987-09-10", "avatar.png", "dave", "Music lover and fitness enthusiast.", 1),
("user6@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Emily", "Taylor", "1991-12-20", "example5.jpg", NULL, "Animal lover and aspiring writer.", 0),
("user7@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "James", "Davis", "1983-06-25", "avatar.png", "jd", "Sports enthusiast and gamer.", 1),
("user8@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Jessica", "Wilson", "1994-08-02", "avatar.png", "jess", "Travel junkie and photography enthusiast.", 0),
("user9@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Matthew", "Anderson", "1980-02-14", "example7.jpg", NULL, "Coffee aficionado and tech geek.", 1),
("user10@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Lauren", "Martinez", "1989-10-18", "avatar.png", "laur", "Nature lover and yoga enthusiast.", 0),
("user11@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Justin", "Moore", "1992-04-08", "example9.jpg", "justinm", "Tech enthusiast and aspiring entrepreneur.", 1),
("user12@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Natalie", "Clark", "1988-09-17", "example2.jpg", NULL, "Music lover and concertgoer.", 0),
("user13@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Brandon", "Hill", "1995-11-25", "example8.jpg", "bhill", "Fitness enthusiast and outdoor adventurer.", 1),
("user14@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Julia", "Young", "1983-03-12", "avatar.png", "jules", "Travel addict and foodie.", 0),
("user15@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Kevin", "Adams", "1986-06-30", "avatar.png", "kev", "Book lover and amateur photographer.", 1),
("user16@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Rachel", "Wright", "1990-12-07", "example4.jpg", NULL, "Coffee addict and Netflix binge-watcher.", 0),
("user17@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Ryan", "Ross", "1981-08-14", "avatar.png", "ryanr", "Sports fanatic and outdoor enthusiast.", 1),
("user18@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Amber", "Gonzalez", "1993-01-28", "avatar.png", "amb", "Tech geek and coffee lover.", 0),
("user19@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Tyler", "Parker", "1987-07-05", "example10.jpg", NULL, "Music lover and aspiring chef.", 1),
("user20@example.com", "$2a$10$dJYBxUfIT8Y5uo9xDETfJO9aDzXhE6qcU4rFdHEhQMJ1V1YaPve2S", "Samantha", "Evans", "1996-09-22", "avatar.png", "sam", "Art enthusiast and nature lover.", 0);

INSERT INTO follows (user_id, follow_id, confirmed)
VALUES 
(1, 3, 1), (2, 4, 0), (5, 7, 2), (6, 1, 1), (9, 8, 0), (10, 11, 1), (12, 5, 2), (14, 19, 0), (15, 13, 1), (16, 18, 2),
(17, 20, 0), (3, 10, 1), (4, 16, 0), (7, 2, 2), (8, 6, 1), (11, 14, 0), (13, 9, 1), (18, 17, 2), (19, 15, 0), (20, 12, 1),
(3, 5, 0), (4, 1, 1), (7, 6, 2), (8, 11, 0), (1, 2, 1), (1, 4, 2), (15, 20, 0), (17, 14, 1), (19, 18, 0), (20, 3, 1),
(1, 9, 2), (2, 15, 0), (5, 17, 1), (6, 20, 2), (9, 13, 0), (10, 19, 1), (12, 2, 2), (14, 4, 0), (16, 7, 1), (18, 8, 2);

INSERT INTO groups (title, description, admin_id)
VALUES
    ("TECH MAVENS", "A group for tech enthusiasts discussing the latest gadgets and innovations.", 9),
    ("ATHLETIC ALL-STARS", "Bringing together athletes from different sports to share training tips and success stories.", 14),
    ("PARENTS' HUB", "A supportive community for parents to exchange parenting advice and family-friendly activities.", 6),
    ("ARTISTIC CREATIONS", "A space for artists to showcase their work, share techniques, and collaborate on projects.", 11),
    ("FREELANCERS UNITED", "Connecting freelancers across industries to share gigs, resources, and insights.", 18),
    ("BOOK LOVERS CLUB", "Dedicated to discussing favorite books, recommending reads, and hosting literary events.", 2),
    ("CODING GENIUSES", "A group for programmers and developers to exchange code snippets, solve problems, and learn new skills.", 17),
    ("OUTDOOR ADVENTURERS", "Exploring nature together through hiking, camping, and outdoor activities.", 4),
    ("HEALTHY LIVING COMMUNITY", "Supporting each other in adopting healthy habits, sharing fitness routines, and nutritious recipes.", 13),
    ("BUSINESS NETWORK", "Bringing professionals together to network, share industry insights, and collaborate on projects.", 8);

INSERT INTO group_user (group_id, user_id, confirmed)
VALUES
    (1, 9, 2), (2, 14, 2), (3, 6, 2), (4, 11, 2), (5, 18, 2), (6, 2, 2), (7, 17, 2), (8, 4, 2), (9, 13, 2), (10, 8, 2), (1, 1, 2),
    (1, 5, 1), (2, 19, 0), (3, 3, 2), (4, 7, 1), (5, 12, 2), (6, 16, 0), (7, 10, 2), (8, 15, 1), (9, 1, 2), (10, 20, 0), (2, 1, 1),
    (1, 18, 1), (2, 2, 0), (3, 14, 2), (4, 8, 1), (5, 6, 2), (6, 13, 0), (7, 9, 2), (8, 17, 1), (9, 4, 2), (10, 11, 0), (3, 1, 0);

INSERT INTO posts (user_id, privacy, text, image, timestamp) 
VALUES
    (9, 3, "Excited to discuss the latest tech innovations with fellow mavens!", "", "2024-03-16 19:46:12"),
    (5, 3, "Looking for recommendations on the best programming languages for beginners.", "", "2023-08-29 08:15:00"),
    (14, 3, "Just finished a morning run! Who else is part of the All-Stars team?", "", "2022-11-05 13:30:45"),
    (19, 3, "Share your favorite workout routines here, All-Stars!", "", "2022-12-20 21:55:30"),
    (14, 3, "Seeking advice on parenting strategies from fellow members of the Parents'' Hub.", "", "2023-05-10 10:10:10"),
    (3, 3, "Let''s share some heartwarming parenting stories, Parents'' Hub!", "", "2023-07-18 17:20:25"),
    (11, 3, "Feeling inspired to create some art today! Who''s with me, Artistic Creations?", "", "2022-09-14 11:11:11"),
    (7, 3, "Share your latest masterpieces here, Artistic Creations!", "", "2023-01-02 15:40:55"),
    (18, 3, "Looking for some freelance gigs! Any opportunities, Freelancers United?", "", "2022-10-30 08:30:20"),
    (6, 3, "Let''s discuss tips and tricks for successful freelancing, Freelancers United!", "", "2023-03-27 14:12:35"),
    (16, 3, "Starting a new book today! Any recommendations from the Book Lovers Club?", "", "2023-02-18 09:25:00"),
    (13, 3, "Discussing the latest plot twists with fellow bookworms in the Book Lovers Club!", "", "2023-09-03 19:00:45"),
    (9, 3, "Excited to dive deep into some coding challenges today! Who''s up for it, Coding Geniuses?", "", "2023-12-10 16:55:30"),
    (17, 3, "Discussing the intricacies of algorithms and data structures with fellow Coding Geniuses!", "", "2022-12-05 12:45:15"),
    (17, 3, "Planning my next hiking trip! Any recommendations, Outdoor Adventurers?", "", "2023-06-20 07:40:00"),
    (4, 3, "Share your most thrilling outdoor adventures here, Outdoor Adventurers!", "", "2022-11-28 18:20:10"),
    (1, 3, "Starting a new fitness challenge! Who wants to join me, Healthy Living Community?", "", "2022-10-15 14:30:45"),
    (13, 3, "Discussing nutrition and wellness tips with fellow members of the Healthy Living Community!", "", "2023-08-07 20:15:30"),
    (8, 3, "Seeking advice on scaling my business. Any insights, Business Network?", "", "2023-04-25 10:05:20"),
    (20, 3, "Let''s share success stories and challenges in the Business Network!", "", "2022-12-30 13:50:55");  

INSERT INTO group_post (post_id, group_id)
VALUES
    (1, 1), (2, 1), (3, 2), (4, 2), (5, 3), (6, 3), (7, 4), (8, 4), (9, 5), (10,5),
    (11, 6), (12, 6), (13, 7), (14, 7), (15, 8), (16, 8), (17, 9), (18, 9), (19, 10), (20, 10);

INSERT INTO comments (post_id, user_id, text, image, timestamp) 
VALUES
    (1, 18, "I'm excited too! Have you heard about the new AI breakthroughs?", "", "2024-03-16 20:30:00"),
    (1, 5, "Looking forward to discussing the latest trends!", "", "2024-03-16 21:15:00"),
    (2, 18, "Python is a great language for beginners!", "", "2023-08-29 09:00:00"),
    (3, 2, "I just finished a run too! Feels great, right?", "", "2022-11-05 14:00:00"),
    (3, 19, "Count me in! All-Star for life!", "", "2022-11-05 14:45:00"),
    (4, 14, "I love sharing workout routines! Here's mine...", "", "2022-12-20 22:30:00"),
    (5, 6, "Parenting is tough but rewarding. Happy to share advice!", "", "2023-05-10 10:45:00"),
    (5, 14, "Let's create a supportive community for parents!", "", "2023-05-10 11:30:00"),
    (6, 3, "Parenting stories always warm my heart. Thanks for sharing!", "", "2023-07-18 18:00:00"),
    (7, 11, "Art is life! What medium do you work with?", "", "2022-09-14 12:00:00"),
    (7, 8, "Feeling inspired as well! Here's my latest piece...", "", "2022-09-14 12:45:00"),
    (8, 7, "I'm amazed by everyone's talent here!", "", "2023-01-02 16:30:00"),
    (9, 18, "I'm always on the lookout for freelancing opportunities!", "", "2022-10-30 09:15:00"),
    (9, 12, "Let's share resources and help each other succeed!", "", "2022-10-30 10:00:00"),
    (10, 6, "Freelancing can be tough, but it's worth it in the end!", "", "2023-03-27 15:30:00"),
    (11, 2, "I just finished a great book! Have you read it?", "", "2023-02-18 10:00:00"),
    (11, 16, "I'm always looking for new recommendations!", "", "2023-02-18 10:45:00"),
    (12, 13, "The plot twists lately have been mind-blowing!", "", "2023-09-03 20:15:00"),
    (13, 17, "Let's tackle some coding challenges together!", "", "2023-12-10 17:30:00"),
    (13, 10, "I love discussing algorithms! What's your favorite topic?", "", "2023-12-10 18:15:00"),
    (14, 10, "I'm always up for a coding discussion!", "", "2022-12-05 13:00:00"),
    (15, 15, "Hiking is one of my favorite activities!", "", "2023-06-20 08:00:00"),
    (15, 17, "Let's plan a group hike soon!", "", "2023-06-20 09:00:00"),
    (16, 15, "I went skydiving last week - an unforgettable experience!", "", "2022-11-28 19:30:00"),
    (17, 13, "Count me in for the fitness challenge!", "", "2022-10-15 15:15:00"),
    (17, 4, "Let's motivate each other to stay healthy!", "", "2022-10-15 16:00:00"),
    (18, 13, "Nutrition is key to a healthy lifestyle. Share your tips!", "", "2023-08-07 21:45:00"),
    (19, 11, "Scaling a business requires careful planning. Let's discuss strategies!", "", "2023-04-25 10:45:00"),
    (19, 8, "Success stories are always inspiring. Thanks for sharing!", "", "2023-04-25 11:30:00"),
    (20, 20, "Overcoming challenges in business is part of the journey!", "", "2022-12-30 14:15:00");

INSERT INTO events (group_id, user_id, title, description, date) 
VALUES
    (1, 9, "Tech Summit 2024", "Join us for a summit on the latest technology trends and innovations.", "2024-03-13"),
    (2, 14, "All-Stars Marathon", "Get ready for the annual All-Stars Marathon! Lace up your shoes and hit the track.", "2024-03-13"),
    (3, 6, "Parenting Workshop", "A workshop for parents by parents. Share your experiences and learn from others.", "2024-03-13"),
    (4, 11, "Art Exhibition", "Discover the beauty of art at our community exhibition. Bring your friends and family!", "2024-03-13"),
    (5, 18, "Freelancer's Meetup", "Connect with fellow freelancers and explore new opportunities at our monthly meetup.", "2024-03-13"),
    (6, 2, "Book Discussion", "Join us for an in-depth discussion of this month's selected book.", "2024-03-13"),
    (7, 17, "Coding Challenge Day", "Test your coding skills with our series of challenges. Prizes await the winners!", "2024-03-13"),
    (8, 4, "Nature Trail Exploration", "Embark on an adventure through scenic nature trails. Don't forget your camera!", "2024-03-13"),
    (9, 13, "Yoga in the Park", "Unwind and rejuvenate with a relaxing yoga session amidst nature's tranquility.", "2024-03-13"),
    (10, 8, "Business Networking Mixer", "Expand your professional network and explore collaboration opportunities over drinks.", "2024-03-13");

INSERT INTO event_users (event_id, user_id, going)
VALUES
    (1, 9, 1), (2, 14, 1), (3, 6, 0), (4, 11, 0), (5, 18, 1), (6, 2, 1), (7, 17, 0), (8, 4, 1), (9, 13, 1), (10, 8, 0);

INSERT INTO posts (user_id, privacy, text, image)
VALUES
    (1, 0, "Just had an amazing day out with friends! It's always great to spend quality time with the people who matter most.", ""),
    (2, 1, "Starting a new chapter in life can be both exhilarating and nerve-wracking. Here's to embracing change and all the possibilities it brings.", ""),
    (3, 0, "There's nothing quite like curling up with a good book on a lazy afternoon. What's your favorite book to unwind with?", ""),
    (4, 1, "Traveling opens up a world of opportunities for exploration and discovery. Where's your next adventure taking you?", ""),
    (5, 0, "Experimenting with recipes in the kitchen is my idea of a fun time. Who else loves getting creative with cooking?", ""),
    (6, 1, "Gratitude turns what we have into enough. Take a moment today to appreciate the little things that make life beautiful.", ""),
    (7, 0, "Popcorn ✔️ Cozy blanket ✔️ Favorite movie ✔️ Looks like it's going to be a perfect movie night in!", ""),
    (8, 1, "Finding balance between work and play is key to living a fulfilling life. How do you maintain harmony in your daily routine?", ""),
    (9, 0, "Sundays are made for relaxation and rejuvenation. How are you spending your day unwinding?", ""),
    (10, 1, "The journey of a thousand miles begins with a single step. What steps are you taking today towards your goals?", ""),
    (11, 0, "A cup of coffee in the morning sets the tone for the day ahead. Who else needs their daily caffeine fix to kickstart the morning?", ""),
    (12, 1, "Celebrating milestones reminds us of how far we've come and motivates us to keep moving forward. What milestones are you celebrating today?", ""),
    (13, 0, "Taking a leisurely stroll through nature is the perfect way to clear the mind and appreciate the beauty around us.", ""),
    (14, 1, "Weekend getaways are a great way to escape the hustle and bustle of everyday life and recharge the soul. Where's your favorite weekend retreat?", ""),
    (15, 0, "Channeling my creativity into DIY projects brings me so much joy. What DIY projects are you currently working on?", ""),
    (16, 1, "Reflection allows us to see how far we've come and where we want to go. Take a moment today to reflect on your journey and express gratitude for the blessings in your life.", ""),
    (17, 0, "Rainy days call for cozy blankets, hot cocoa, and good company. How do you like to spend rainy days?", ""),
    (18, 1, "Cherishing moments with loved ones creates memories that last a lifetime. What's your favorite memory with someone special?", ""),
    (19, 0, "Painting is my favorite form of self-expression. There's something truly therapeutic about putting brush to canvas and letting creativity flow.", ""),
    (20, 1, "Dream big, set goals, and chase them relentlessly. Your dreams are within reach if you're willing to put in the effort and stay focused.", "");

INSERT INTO chat (to_id, from_id, message, timestamp)
VALUES 
    (1, 2, "Hey there! How's it going?", "2024-01-01 00:00:00"),
    (2, 1, "Hey! I'm doing well, thanks for asking.", "2024-01-01 00:00:01"),
    (1, 2, "That's great to hear!", "2024-01-01 00:00:02"),
    (2, 1, "Yeah! How about you?", "2024-01-01 00:00:03"),
    (1, 2, "I'm doing okay, just a bit tired.", "2024-01-01 00:00:04"),
    (2, 1, "Take some rest, it's important.", "2024-01-01 00:00:05"),
    (1, 2, "Yeah, I will. Thanks.", "2024-01-01 00:00:06"),
    (1, 2, "Did you watch the latest movie?", "2024-01-01 00:00:07"),
    (2, 1, "No, not yet. Was it good?", "2024-01-01 00:00:08"),
    (1, 2, "It was amazing! You should definitely watch it.", "2024-01-01 00:00:09"),
    (2, 1, "I'll check it out this weekend.", "2024-01-01 00:00:10"),
    (1, 2, "Sure, let me know how you find it.", "2024-01-01 00:00:11"),
    (2, 1, "Will do!", "2024-01-01 00:00:12"),
    (1, 2, "By the way, have you tried that new restaurant?", "2024-01-01 00:00:13"),
    (2, 1, "Not yet. How was it?", "2024-01-01 00:00:14"),
    (1, 2, "It was fantastic! You should try their pasta.", "2024-01-01 00:00:15"),
    (2, 1, "I'll make sure to visit soon.", "2024-01-01 00:00:16"),
    (1, 2, "Great! Let me know how you like it.", "2024-01-01 00:00:17"),
    (2, 1, "Definitely!", "2024-01-01 00:00:18"),
    (1, 2, "I need some advice on buying a new phone.", "2024-01-01 00:00:19"),
    (2, 1, "Sure, I can help. What are you looking for?", "2024-01-01 00:00:20"),
    (1, 2, "Something with a good camera and long battery life.", "2024-01-01 00:00:21"),
    (2, 1, "I have a few options in mind. I'll send them to you.", "2024-01-01 00:00:22"),
    (1, 2, "Thanks, I appreciate it.", "2024-01-01 00:00:23"),
    (4, 1, "Hey, how's your day going?", "2024-01-01 00:00:24"),
    (1, 4, "Hi! It's been busy, but good overall. How about you?", "2024-01-01 00:00:25"),
    (4, 1, "I'm just relaxing after work. Do anything interesting today?", "2024-01-01 00:00:26"),
    (1, 4, "Not really, just the usual. How about you?", "2024-01-01 00:00:27"),
    (4, 1, "I went for a walk in the park. It was refreshing.", "2024-01-01 00:00:28"),
    (1, 4, "That sounds nice. I should do that sometime.", "2024-01-01 00:00:29"),
    (4, 1, "Yeah, it's a good way to clear your mind.", "2024-01-01 00:00:30"),
    (4, 1, "By the way, have you seen the new exhibition at the museum?", "2024-01-01 00:00:31"),
    (1, 4, "No, not yet. Is it worth checking out?", "2024-01-01 00:00:32"),
    (4, 1, "Definitely! The artworks are amazing.", "2024-01-01 00:00:33");

INSERT INTO group_chat (group_id, from_id, message, timestamp)
VALUES 
    (1, 1, "Hey everyone, how's it going?", "2024-01-01 00:00:00"),
    (1, 9, "I'm doing well, thanks!", "2024-01-01 00:00:01"),
    (1, 5, "Hey guys! What's up?", "2024-01-01 00:00:02"),
    (1, 18, "Not much, just chilling.", "2024-01-01 00:00:03"),
    (1, 1, "Has anyone seen the latest episode of that show?", "2024-01-01 00:00:04"),
    (1, 9, "No, I haven't. Was it good?", "2024-01-01 00:00:05"),
    (1, 5, "Yeah, it was awesome!", "2024-01-01 00:00:06"),
    (1, 18, "I'll have to watch it then.", "2024-01-01 00:00:07"),
    (1, 1, "Definitely!", "2024-01-01 00:00:08"),
    (1, 9, "Hey, did you all hear about the new restaurant opening up?", "2024-01-01 00:00:09"),
    (1, 5, "No, tell us more!", "2024-01-01 00:00:10"),
    (1, 18, "What kind of food do they serve?", "2024-01-01 00:00:11"),
    (1, 1, "They serve Italian cuisine.", "2024-01-01 00:00:12"),
    (1, 9, "Sounds delicious!", "2024-01-01 00:00:13"),
    (1, 5, "I'm excited to try it out!", "2024-01-01 00:00:14");
