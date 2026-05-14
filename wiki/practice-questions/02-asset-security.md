---
title: "Practice Questions — Domain 2: Asset Security"
type: practice
domain: 2
tags: [asset-management, data-classification, data-roles, data-states, sanitization, practice]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Practice Questions — Domain 2: Asset Security

## Questions

### Q1: An HR director manages an HR database. The IT department performs nightly backups and applies access controls. If a data breach occurs because the database was misconfigured, who is ultimately accountable?

**Answer:** The HR Director (the data owner).

**Why:** The HR Director is the data owner — the person who directly interacts with the asset most and knows its value. The IT department acts as the data custodian with technical responsibility, but accountability for the protection of the asset cannot be delegated. The owner remains accountable even when responsibility for day-to-day operations has been handed to IT.

*Subtopic: 2.3 — Data roles and ownership*  
*Concept: [Data Roles](../concepts/data-roles.md)*

---

### Q2: A manager asks IT to implement access controls on a newly created dataset. Who should determine the classification level and approve access to that dataset?

**Answer:** The data owner (the manager, or whoever is accountable for the asset's business value).

**Why:** Classification is driven by the data owner, not by IT (the custodian) and not by the security team. The owner understands the asset's value and is accountable for classifying it and approving access. IT implements what the owner specifies — IT's role is custodial.

*Subtopic: 2.3 — Data classification roles*  
*Concept: [Data Classification](../concepts/data-classification.md)*

---

### Q3: An organization uses "Top Secret" and "Confidential" as classification labels, as does a partner organization. A security analyst assumes that assets labeled "Confidential" at both organizations are protected to the same standard. What is wrong with this assumption?

**Answer:** Different organizations may use identical classification labels but assign them completely different values and protection baselines.

**Why:** Classification is a system of classes; the same label can represent very different asset values depending on how each organization has defined its classification system. Security teams must educate users about what each label means *within their own organization* and not assume cross-organization consistency of labels.

*Subtopic: 2.1 — Classification vs. categorization*  
*Concept: [Data Classification](../concepts/data-classification.md)*

---

### Q4: Which of the following best describes the difference between security labeling and security marking, per NIST SP 800-53A?

A) Labeling is applied by humans; marking is applied by systems  
B) Labeling associates security attributes with internal data structures for system-based enforcement; marking associates security attributes with objects in human-readable form for process-based enforcement  
C) Labeling applies to physical media; marking applies to digital assets  
D) Labeling and marking are interchangeable terms

**Answer:** B

**Why:** Per NIST SP 800-53A (cited in destination-cissp §2.1.4): labeling is machine/system-readable and enables automated policy enforcement; marking is human-readable and enables process-based (human) enforcement. Common exam trap: many candidates reverse these.

*Subtopic: 2.1 — Labeling and marking*  
*Concept: [Data Classification](../concepts/data-classification.md)*

---

### Q5: A security engineer needs to sanitize a hard drive containing top-secret data before the drive is disposed of. The engineer has time and resources available. Which method provides the highest assurance that data cannot be recovered?

**Answer:** Physical destruction — specifically incineration.

**Why:** Per NIST SP 800-88 Rev. 1, the sanitization categories from most to least effective are Destroy > Purge > Clear. Incineration is the most effective form of destruction, rendering the media to molten material and providing the highest assurance. Shredding and drilling are also in the Destroy category but not foolproof (skilled attackers may recover data from platters). Overwriting is only Clear; degaussing is Purge/Destroy boundary.

*Subtopic: 2.4 — Data destruction*  
*Standard: [NIST SP 800-88](../standards/nist-sp-800-88.md)*

---

### Q6: An organization is decommissioning a server hosted by a cloud provider. A security engineer recommends crypto shredding. Under what condition does this qualify as "purging" rather than merely "clearing"?

**Answer:** Crypto shredding qualifies as purging when all copies of the encryption key are provably destroyed and the encryption algorithm cannot be feasibly brute-forced.

**Why:** Crypto shredding (encrypting data then destroying the key) sits between Purge and Clear in the sanitization hierarchy. If the key is irretrievable and the algorithm is sound, the data is unrecoverable → Purge. If there is any chance the key could be found (backup key, weak algorithm) → effectively only Clear. Physical destruction of cloud media is typically infeasible, making crypto shredding the recommended cloud alternative.

*Subtopic: 2.4 — Data remanence, cloud data destruction*  
*Concept: [Data Security Controls](../concepts/data-security-controls.md)*

---

### Q7: Which encryption approach protects data while it is being actively computed on, without requiring the data to be decrypted?

**Answer:** Homomorphic encryption.

**Why:** Homomorphic encryption allows mathematical computations to be performed on ciphertext, producing an encrypted result that — when decrypted — matches the result of operating on the plaintext. This protects data in use. It is the only method that addresses the unique security challenge of data in active computational processing. RBAC and DLP also protect data in use but do not encrypt the data during computation.

*Subtopic: 2.6 — Protecting data in use*  
*Concept: [Data States](../concepts/data-states.md)*

---

### Q8: An organization classifies a production queueing system as "high availability required" even though the data it processes is not sensitive. What does this illustrate about modern asset classification?

**Answer:** Classification should be based on all three CIA dimensions (confidentiality, integrity, and availability), not just confidentiality/sensitivity.

**Why:** Traditional classification often focuses only on confidentiality (sensitivity labels like "Secret" or "Confidential"). Modern guidance — and the destination-cissp source — recommends assigning separate classification scores for confidentiality (sensitivity), integrity (accuracy), and availability (criticality). A highly available but non-sensitive system legitimately warrants a high classification on the availability axis.

*Subtopic: 2.1 — Classification based on CIA*  
*Concept: [Data Classification](../concepts/data-classification.md)*

---

### Q9: End-to-end encryption is applied to data traversing a network. A network analyst captures packets at an intermediate router. What information is visible to the analyst?

**Answer:** The routing information (source and destination IP addresses). The data payload is encrypted and not readable.

**Why:** End-to-end encryption encrypts the data portion of the packet at the source node and keeps it encrypted through every intermediate node. The data is only decrypted at the destination. However, the packet headers — including IP addresses — remain in plaintext. This means end-to-end encryption does not provide anonymity, and traffic analysis attacks on routing metadata are possible. Contrast with onion routing, which hides source and destination.

*Subtopic: 2.6 — Protecting data in transit*  
*Concept: [Data States](../concepts/data-states.md)*

---

### Q10: A developer replaces real customer Social Security Numbers (SSNs) in a test database with randomly generated 9-digit numbers that bear no relationship to the original values, and stores no mapping table. Which obfuscation technique is this, and is re-identification possible?

**Answer:** Fabrication (or full anonymization of the field). Re-identification is not possible because no mapping exists between the fake values and the originals.

**Why:** Fabrication replaces real data with realistic-looking fake data; with no mapping table, the original data cannot be reconstructed — this is the strongest obfuscation for test environments. Contrast with tokenization (a mapping table is kept in a token vault — re-identification IS possible) and pseudonymization (mapping exists somewhere). If no mapping is retained, fabrication achieves the same protection as anonymization for that field.

*Subtopic: 2.6 — Information obfuscation*  
*Concept: [Information Obfuscation](../concepts/information-obfuscation.md)*
