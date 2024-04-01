using FreightCMS.Models;
using FreightCMS.Locations.Models;

namespace FreightCMS.Parties.Models;

public class Party
{
    /// <summary>
    /// Unique Identifier for the Party that can be database or application generated.
    /// </summary>
    public required string Id { get; set; }
    /// <summary>
    /// Name of the Party that is used for display purposes.
    /// </summary>
    public required string Name { get; set; }
    /// <summary>
    /// Party Type can be Shipper, Consignee, Carrier, Broker, etc.
    /// </summary>
    public required string PartyType { get; set; }    
}

public class Person : Party
{
    /// <summary>
    /// Tax Identification Number of the Company.
    /// </summary>
    public string? TaxId { get; set; }
    public required AddressModel PhysicalAddress { get; set; }
    public required AddressModel MailingAddress { get; set; }
    public required Contact ContactInfo { get; set; }
}

public class Company : Party
{
    /// <summary>
    /// Tax Identification Number of the Company.
    /// </summary>
    public string? TaxId { get; set; }

    /// <summary>
    /// Doing Business As Name of the Company.
    /// </summary>
    public required string DBA { get; set; }
    public required AddressModel PhysicalAddress { get; set; }
    public required AddressModel MailingAddress { get; set; }
    /// <summary>
    /// List of Contacts associated with the Company.
    /// </summary>
    public required ICollection<Contact> Contacts { get; set; } = [];
}