# User Service Agents Guide

This service owns core user social data.

## Scope

- user-facing base profile data such as name, bio, avatar, and similar social attributes
- APIs and persistence for that source-of-truth data

## Working Rules

- Keep source-of-truth ownership here; aggregation or presentation concerns belong elsewhere.
- Review impact on `profile` and client modules when user fields change.
- Be careful with media, identity, and schema changes because they tend to ripple through multiple consumers.
