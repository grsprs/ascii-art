# Security Policy

## Reporting Security Issues

Found a security problem? Here's how to report it safely.

### Contact Method

**Email**: sp.nikoloudakis@gmail.com  
**Response Time**: Reports acknowledged within 48 hours  
**Subject**: Please use [SECURITY] in the subject line

### What to Include

1. **Description** of the vulnerability
2. **Steps to reproduce** with minimal example
3. **Impact assessment** (confidentiality, integrity, availability)
4. **Suggested fix** (if known)
5. **Your contact information** for follow-up

### Example Report Format

```
Subject: [SECURITY] Buffer overflow in banner parser

Description:
The banner file parser does not validate line length, allowing 
potential buffer overflow with crafted banner files.

Steps to Reproduce:
1. Create banner file with line > 10KB
2. Run: ascii-art -banner=malicious "test"
3. Observe memory corruption

Impact:
- Potential code execution
- Denial of service
- Memory corruption

Environment:
- OS: Linux x64
- Go version: 1.20.1
- ascii-art version: 1.2.3
```

## Security Measures

### Input Validation

- **Size limits**: Input capped at 1MB by default
- **Character validation**: Only printable ASCII (32-126)
- **Banner file validation**: Format and size checks
- **Path validation**: Prevent directory traversal

### Memory Safety

- **Bounds checking**: All array/slice access validated
- **Resource limits**: Memory usage monitoring
- **Timeout protection**: Prevent infinite loops
- **Safe string operations**: Use Go's safe string functions

### Output Safety

- **No code injection**: Output is plain text only
- **Encoding safety**: UTF-8 compliant output
- **File permissions**: Restrictive output file permissions
- **Path sanitization**: Safe output file handling

## Threat Model

### Assets

- **User input data**: Text to be rendered
- **Banner files**: ASCII art templates
- **Output**: Generated ASCII art
- **System resources**: Memory, CPU, disk

### Threats

1. **Malicious input**: Crafted strings causing DoS
2. **Banner file tampering**: Modified templates
3. **Resource exhaustion**: Memory/CPU abuse
4. **Path traversal**: Unauthorized file access
5. **Code injection**: If output used in other systems

### Mitigations

- Input size and character validation
- Banner file integrity checks
- Resource usage monitoring
- Safe file operations
- Output sanitization guidance

## Vulnerability Disclosure

### Timeline

- **T+0**: Report received, acknowledgment sent
- **T+7 days**: Initial assessment completed
- **T+30 days**: Fix developed and tested
- **T+45 days**: Security update released
- **T+90 days**: Public disclosure (if resolved)

### Severity Classification

- **Critical**: Remote code execution, privilege escalation
- **High**: Local code execution, significant DoS
- **Medium**: Information disclosure, limited DoS
- **Low**: Minor information leakage

## Security Artifacts

### CI/CD Security

- **SBOM Generation**: Software Bill of Materials for each release
- **Dependency Scanning**: Automated vulnerability detection
- **Static Analysis**: Security-focused code analysis
- **License Scanning**: Open source license compliance

### Artifacts Available

- `SECURITY_REPORTS/sbom-cyclonedx.json` - Software Bill of Materials
- `SECURITY_REPORTS/dependency-scan.json` - Vulnerability report
- `SECURITY_REPORTS/static-analysis.json` - Code analysis results
- `SECURITY_REPORTS/license-scan.json` - License compliance

## Security Testing

### For Security Researchers

**Test Environment Setup**:
```bash
# Safe testing environment
docker run -it --rm golang:1.20
git clone https://github.com/grsprs/ascii-art.git
cd ascii-art
go build ./cmd/ascii-art
```

**Test Cases to Consider**:
- Large input strings (>1MB)
- Malformed banner files
- Special characters and escape sequences
- Concurrent usage patterns
- File system edge cases

### Responsible Disclosure

- Test only in isolated environments
- Do not access production systems
- Do not exfiltrate data
- Report findings promptly
- Allow reasonable time for fixes

## Security Updates

### Notification Channels

- GitHub Security Advisories
- Release notes with security tags
- Email notifications (if subscribed)

### Update Process

1. Security fix developed
2. Automated testing completed
3. Security advisory published
4. Release with security tag
5. Notification sent to users

## Compliance

### Standards Alignment

- **OWASP**: Following secure coding practices
- **CWE**: Common Weakness Enumeration awareness
- **CVE**: Common Vulnerabilities and Exposures tracking

### Enterprise Security

- **SBOM**: Software Bill of Materials provided
- **SCA**: Software Composition Analysis reports
- **SAST**: Static Application Security Testing
- **Dependency Tracking**: Known vulnerability monitoring

## Contact

For non-security issues, use GitHub issues.  
For security concerns, use the contact method above.

**Security Contact**: sp.nikoloudakis@gmail.com  
**Project Author**: Spiros Nikoloudakis

---

**Last updated: November 25, 2024**