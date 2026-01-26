Astana IT University. 
2nd year of Software Engineering. 
Advanced Programming 1 (GoLang) course. 
Students: Alikhan Mukhamed-Rakhym, Daniyal Adilbekuly, Alexander Tarulin.
Group: SE-2407 & SE-2408

Assignment 3 Milestone 1 Report – “LifeFlow” Project.

Project topic: Task & Life Management System







1. Project Proposal

1.1 Project Overview
Purpose of the project
LifeFlow is a web-based task and life management system designed to help users organize daily activities, manage tasks, and track progress in a simple and efficient way. The project is developed as part of the Advanced Programming 1 course and focuses on applying backend development principles using the Go programming language.

1.2 Project Relevance 
Modern users face difficulties managing daily responsibilities, deadlines, and personal goals. While many task management tools exist, they are often overloaded with features or difficult to use. LifeFlow aims to provide a lightweight and intuitive solution that focuses on essential task and life organization features.

1.3 Competitor Analysis
Examples:
•	Todoist
•	Trello
•	Microsoft To Do
Existing solutions offer advanced features such as team collaboration and integrations. However, LifeFlow focuses on simplicity, individual productivity, and educational purposes, allowing users to manage tasks without unnecessary complexity. Also our project will have unique features which will solve main problems and drawbacks of existing competitors.

1.4 Target Audience
The target audience of LifeFlow includes:
- learners
- teachers
- students
- young professionals
- any person who wants to develop and improve their skills/life!
- or any individuals who want to organize their daily tasks and personal goals using a simple web application.

1.5 Planned Features
•	User registration and login
•	Task creation, editing, and deletion (CRUD)
•	Task status management (To Do / In Progress / Done)
•	Habit tracking (optional)
•	Simple dashboard
•	Server-side rendered HTML pages
•	Improved and convenient features which are better than competitors













2.  Architecture & Design
2.1 Architectural Style
LifeFlow follows a monolithic architecture, where all application components are implemented within a single Go application. This approach simplifies development, deployment, and maintenance, making it suitable for educational projects and small-scale systems.
2.2 High-Level System Architecture
 







2.3 Use Case Diagram
 
2.4 ERD (Entity Relationship Diagram)
 

2.5 UML Class Diagram
 

2.6 Project Structure Explanation (Data flow, modules and responsibilities)
The project is structured following Go best practices. The cmd directory contains the application entry point, while the internal directory contains application logic such as handlers, services, repositories, and models. This structure ensures separation of concerns and maintainability.





3.  Project Plan
3.1 Gantt Chart (Weeks 7–10)
 
Week 7
•	Project planning and requirement analysis 
•	GitHub repository setup and project initialization 
•	System architecture and design preparation 
•	Basic CRUD implementation
•	Basic Database integration
•	Basic backend and server setup
•	(Maybe) The first well designed front-end page
•	Creating the first working version of the project
Week 8
•	Refinement of backend architecture and code structure
•	Database integration improvement and optimization
•	Authentication and authorization implementation
•	Additional CRUD functionality and business logic
•	Frontend page development and UI improvements
•	Integration of backend and frontend components
Week 9
•	Final feature completion and system stabilization
•	Middleware implementation (logging and authentication)
•	Bug fixing and code cleanup
•	Testing of application functionality
•	Performance and usability improvements
•	Presentation preparation








3.2 Team Responsibilities

Team Member	Responsibilities
Alikhan Mukhamed-Rakhym 	Backend architecture, system integration, and overall coordination, DB integration, Front-end integration and improvement, Coordinating development between team members, Ensuring project alignment with course requirements
Daniyal Adilbekuly	Task entity backend functionality including HTTP routing, handlers, service layer, repository pattern, and in-memory CRUD logic. Front-end improvement and integration of front-end features
Alexander Tarulin	Mainly responsible for Database integration and development, adding more features through backend and improving overall project functionality.



4.  Repository Setup

4.1 Version Control
 
 
GitHub is used for version control. The repository contains a main branch and separate feature branches for each team member.

4.2 Collaboration
Each team member contributes through individual branches and commits. Pull requests are used to merge features into the main branch.

4.3 Build & Run Instructions
go run  .

Conclusion
This milestone demonstrates the planning, architecture, and design decisions for the LifeFlow project. The provided documentation, diagrams, and project plan form a solid foundation for further implementation in the final project stage.

