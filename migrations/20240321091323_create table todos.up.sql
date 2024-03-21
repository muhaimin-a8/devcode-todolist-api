CREATE TABLE todos (
  todo_id INT PRIMARY KEY AUTO_INCREMENT,
  activity_group_id INT NOT NULL,
  FOREIGN KEY (activity_group_id) REFERENCES activities(activity_id),
  title VARCHAR(255) NOT NULL,
  priority VARCHAR(20) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
