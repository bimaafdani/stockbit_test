   
    SELECT user.ID, user.Username, 
    CASE
        WHEN user1.Username IS NULL THEN 'None'
        ELSE user1.Username
    END AS 'ParentUserName'
    FROM User user 
    LEFT JOIN User user1 ON user.Parent = user1.ID
    ORDER BY user.ID
