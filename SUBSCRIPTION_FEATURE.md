# Subscription Feature

## Overview
Simple subscription system with two tiers: Free and Premium.

## User Roles
- `free` - Default role for new users
- `premium` - Upgraded users with full access

## Feature Restrictions

### Free Users
1. **Development Page**: Cannot use folders (all tasks go to "General" folder)
2. **Notes Page**: Limited to 10 notes maximum
3. **Health Page**: Cannot view sleep charts (only see records)

### Premium Users
1. **Development Page**: Full folder functionality
2. **Notes Page**: Unlimited notes
3. **Health Page**: Full access to charts and analytics

## How to Upgrade
1. Go to Profile page
2. Click "Upgrade Now" button
3. User role changes from "free" to "premium"

## Implementation Details

### Database
- Added `role` field to User model (string: "free" or "premium")
- Default value: "free" for new registrations

### Middleware
- `subscription_middleware.go` - Contains role checking logic
- `GetUserRole()` - Helper function to get user's subscription status

### Handlers
- All handlers updated to check user role
- Restrictions enforced server-side
- UI updated to show/hide features based on role

### Templates
- Development: Folder input hidden for free users
- Notes: Shows note count and limit warning
- Health: Charts hidden for free users
- Profile: Shows subscription status and upgrade button
