# Security Configuration Guide

## 🔒 Production Deployment Security Checklist

This document provides security configuration guidelines for deploying New API in production environments.

---

## 1. 🚨 Critical - Must Complete Before Production

### 1.1 Change All Default Passwords

**Database Passwords:**
```yaml
# docker-compose.yml
POSTGRES_PASSWORD: <CHANGE_ME_TO_STRONG_PASSWORD>  # Generate 32+ char random password
MYSQL_ROOT_PASSWORD: <CHANGE_ME_TO_STRONG_PASSWORD>  # If using MySQL
```

**Redis Password:**
```yaml
# docker-compose.yml
redis:
  command: ["redis-server", "--requirepass", "<CHANGE_ME_TO_STRONG_PASSWORD>"]
environment:
  - REDIS_CONN_STRING=redis://:<CHANGE_ME_TO_STRONG_PASSWORD>@redis:6379
```

**Generate Strong Passwords:**
```bash
# Generate a 32-character random password
openssl rand -base64 32

# Or use this command (macOS/Linux)
LC_ALL=C tr -dc 'A-Za-z0-9!@#$%^&*' < /dev/urandom | head -c 32; echo
```

### 1.2 Set Required Secrets

**Multi-node Deployment (Required):**
```bash
# Set SESSION_SECRET for session consistency across nodes
SESSION_SECRET=$(openssl rand -base64 32)
```

**Redis Encryption (Required):**
```bash
# Set CRYPTO_SECRET for Redis data encryption
CRYPTO_SECRET=$(openssl rand -base64 32)
```

**Add to docker-compose.yml:**
```yaml
environment:
  - SESSION_SECRET=<YOUR_GENERATED_SECRET>
  - CRYPTO_SECRET=<YOUR_GENERATED_SECRET>
```

---

## 2. 🌐 CORS Configuration

### 2.1 Restrict Allowed Origins

**Set Environment Variable:**
```bash
# In docker-compose.yml or .env
ALLOWED_ORIGINS=https://yourdomain.com,https://admin.yourdomain.com
```

**For Single Domain:**
```bash
ALLOWED_ORIGINS=https://yourdomain.com
```

**Warning:** If `ALLOWED_ORIGINS` is not set, CORS will allow all origins (insecure for production).

---

## 3. 🔐 HTTPS/TLS Configuration

### 3.1 Enable HTTPS

**Option 1: Use Reverse Proxy (Recommended)**
```nginx
# nginx.conf
server {
    listen 443 ssl http2;
    server_name yourdomain.com;

    ssl_certificate /path/to/fullchain.pem;
    ssl_certificate_key /path/to/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;

    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

**Option 2: Use Let's Encrypt + Certbot**
```bash
# Install certbot
apt-get install certbot python3-certbot-nginx

# Generate certificate
certbot --nginx -d yourdomain.com
```

---

## 4. 🔥 Firewall Configuration

### 4.1 Restrict Database Access

**Only allow localhost connections:**
```yaml
# docker-compose.yml - PostgreSQL
postgres:
  ports:
    # Remove or comment out this line to prevent external access
    # - "5432:5432"
```

### 4.2 Firewall Rules (iptables)

```bash
# Allow HTTPS only
iptables -A INPUT -p tcp --dport 443 -j ACCEPT

# Allow SSH (change 22 to your SSH port)
iptables -A INPUT -p tcp --dport 22 -j ACCEPT

# Block direct access to application port from external
iptables -A INPUT -p tcp --dport 3000 -s 127.0.0.1 -j ACCEPT
iptables -A INPUT -p tcp --dport 3000 -j DROP

# Save rules
iptables-save > /etc/iptables/rules.v4
```

---

## 5. 📊 Rate Limiting Configuration

### 5.1 Adjust Rate Limits

```bash
# Environment variables for rate limiting
# Global Web Rate Limit (requests per minute)
GLOBAL_WEB_RATE_LIMIT_NUM=60
GLOBAL_WEB_RATE_LIMIT_DURATION=60

# Global API Rate Limit (requests per minute)
GLOBAL_API_RATE_LIMIT_NUM=120
GLOBAL_API_RATE_LIMIT_DURATION=60

# Critical Operations Rate Limit (login, register, etc.)
CRITICAL_RATE_LIMIT_NUM=10
CRITICAL_RATE_LIMIT_DURATION=60
```

---

## 6. 🗄️ Database Security

### 6.1 Database Connection Limits

```bash
# Maximum database connections
SQL_MAX_IDLE_CONNS=100
SQL_MAX_OPEN_CONNS=1000
SQL_MAX_LIFETIME=60
```

### 6.2 Enable Error Logging

```bash
ERROR_LOG_ENABLED=true
```

### 6.3 Database Backup Strategy

**PostgreSQL Backup:**
```bash
# Daily backup script
#!/bin/bash
BACKUP_DIR=/backups
DATE=$(date +%Y%m%d_%H%M%S)
docker exec postgres pg_dump -U root new-api > "$BACKUP_DIR/backup_$DATE.sql"

# Keep only last 7 days
find $BACKUP_DIR -name "backup_*.sql" -mtime +7 -delete
```

**Add to crontab:**
```bash
0 2 * * * /path/to/backup-script.sh
```

---

## 7. 🔍 Monitoring & Logging

### 7.1 Enable Comprehensive Logging

```bash
ERROR_LOG_ENABLED=true
LOG_SQL_DSN=<separate_log_database_connection>
```

### 7.2 Log Rotation

```yaml
# docker-compose.yml
services:
  new-api:
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
```

---

## 8. 🛡️ Security Headers

The following security headers are now automatically set (as of the latest update):

- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `X-XSS-Protection: 1; mode=block`
- `Referrer-Policy: strict-origin-when-cross-origin`
- `Strict-Transport-Security: max-age=31536000` (when HTTPS is detected)
- `Content-Security-Policy` (configured for the application)

---

## 9. 🔑 Authentication Security

### 9.1 Enable Two-Factor Authentication

**Recommend to users:**
- Enable 2FA/TOTP in profile settings
- Use Passkey/WebAuthn for passwordless authentication

### 9.2 Password Policy

**Current Requirements:**
- Minimum 8 characters
- Maximum 20 characters
- Bcrypt hashed with default cost (10)

**Recommended Enhanced Policy (requires code change):**
- Minimum 12 characters
- Require uppercase, lowercase, number, and special character

---

## 10. 📦 Dependency Updates

### 10.1 Regular Updates

**Check for updates monthly:**
```bash
# Backend (Go)
go list -u -m all

# Frontend (Bun)
cd web/default
bun update
bun audit
```

### 10.2 Security Audit

```bash
# Run security audit
cd web/default
bun audit

# Review and update vulnerable packages
bun update <package-name>
```

---

## 11. 🚀 Production Environment Variables

### 11.1 Complete Production .env Example

```bash
# === Required Production Settings ===
GIN_MODE=release
NODE_NAME=production-node-1
SESSION_SECRET=<GENERATE_32_CHAR_SECRET>
CRYPTO_SECRET=<GENERATE_32_CHAR_SECRET>

# === CORS Security ===
ALLOWED_ORIGINS=https://yourdomain.com,https://admin.yourdomain.com

# === Database ===
SQL_DSN=postgresql://user:<STRONG_PASSWORD>@postgres:5432/new-api
SQL_MAX_IDLE_CONNS=100
SQL_MAX_OPEN_CONNS=1000
SQL_MAX_LIFETIME=60

# === Redis ===
REDIS_CONN_STRING=redis://:<STRONG_PASSWORD>@redis:6379
MEMORY_CACHE_ENABLED=true
SYNC_FREQUENCY=60

# === Rate Limiting ===
GLOBAL_WEB_RATE_LIMIT_NUM=60
GLOBAL_API_RATE_LIMIT_NUM=120
CRITICAL_RATE_LIMIT_NUM=10

# === Logging ===
ERROR_LOG_ENABLED=true
BATCH_UPDATE_ENABLED=true

# === Timeouts ===
STREAMING_TIMEOUT=300
RELAY_TIMEOUT=0

# === Security ===
TLS_INSECURE_SKIP_VERIFY=false

# === Timezone ===
TZ=Asia/Shanghai
```

---

## 12. ✅ Security Verification Checklist

Before going to production, verify:

- [ ] All default passwords changed
- [ ] `SESSION_SECRET` set (multi-node deployment)
- [ ] `CRYPTO_SECRET` set (Redis encryption)
- [ ] `ALLOWED_ORIGINS` configured (CORS)
- [ ] HTTPS/TLS enabled
- [ ] Firewall rules configured
- [ ] Database external access blocked
- [ ] Rate limiting configured
- [ ] Error logging enabled
- [ ] Database backup strategy in place
- [ ] All dependencies updated (`bun audit` clean)
- [ ] `GIN_MODE=release` set
- [ ] Security headers verified (check response headers)
- [ ] 2FA enabled for admin accounts
- [ ] Log rotation configured

---

## 13. 🔒 Incident Response

### 13.1 If Credentials Are Compromised

1. **Immediately rotate all secrets:**
   ```bash
   # Generate new secrets
   NEW_SESSION_SECRET=$(openssl rand -base64 32)
   NEW_CRYPTO_SECRET=$(openssl rand -base64 32)
   
   # Update docker-compose.yml
   # Restart services
   docker-compose down
   docker-compose up -d
   ```

2. **Revoke all active sessions:**
   - Users will need to log in again

3. **Review access logs:**
   ```bash
   docker logs new-api | grep -i "error\|unauthorized\|failed"
   ```

### 13.2 Security Contact

For security vulnerabilities, please report to:
- GitHub Security Advisories: https://github.com/QuantumNous/new-api/security
- Email: support@quantumnous.com

---

## 14. 📚 Additional Resources

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Mozilla Security Guidelines](https://infosec.mozilla.org/guidelines/)
- [Docker Security Best Practices](https://docs.docker.com/engine/security/)
- [Let's Encrypt Documentation](https://letsencrypt.org/docs/)

---

**Last Updated:** 2026-05-19  
**Version:** 1.0  
**Maintainer:** New API Security Team
