# Federated Data Broker

## Author
Eric Chiquito

## 1 Introduction

The ability to share, discover, and utilize data effectively is becoming an indus-
try’s key competitive advantage. However, many barriers prevent organizations from fully realizing the potential of data sharing. Trust issues, lack of interoperability, data sovereignty concerns, and compliance requirements are just a few of the challenges that hinder the development of a fair and secure data ecosystem. As a result, organizations often operate in isolation, unable to access or benefit from valuable external data sources. A solution to these challenges is essential for building a collaborative, data-rich environment where organizations can connect, share, and utilize data safely and efficiently.

## 2 Project Objective
The project will focus on developing a data broker that facilitates the collection and retrieval of data from various systems based on the work proposed by the International Data Space Association on Metadata Brokers [IDS23], including Enterprise Resource Planning (ERP) systems and real-time transportation services. This broker will enable secure and structured data exchange, allowing Responsible Economic Operators (REOs) to establish agreements that define how their data is shared and with which partners.

The goal is to set up a network consisting of multiple servers running open-source ERP systems such as Odoo [Odo]. Additionally, develop a central data broker that periodically retrieves data from these ERP systems, enabling users to perform queries and data analysis on the collected data.

The work carried out during this project is intended to be integrated into the Demand-Supply Matching (DSM) system developed at LTU, assessing its interoperability with emerging Digital Product Passport (DPP) systems [Cir24].

## 3 How to use
Documentation of the product in `docs/PDB_documentation_v1`

## References
- [IDS23] IDSA. IDS RAM 4: <i>Metadata Broker</i>. https://docs.internationaldataspaces.org/ids-knowledgebase/ids-ram-4/layers-of-the-reference-architecture-model/3-layers-of-the-reference-architecture-model/3_5_0_system_layer/3_5_4_metadata_broker. 2023.
- [Cir24] <i>Cirpass</i>: DPP in a nutshell. Cirpass Project. 2024.
- [Odo] Odoo. https://github.com/odoo/odoo.

