﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions.Azure;
using AutoRest.Go;
using System;
using System.Linq;

namespace AutoRest.Go.Model
{
    /// <summary>
    /// Represents a response from a pageable operation.
    /// </summary>
    internal class PageTypeGo : CompositeTypeGo
    {
        /// <summary>
        /// Creates a new pageable type for the specified response type.
        /// </summary>
        /// <param name="method">The method that will return a pageable type.</param>
        public PageTypeGo(MethodGo method) : base(CodeNamerGo.Instance.GetPageTypeName(method))
        {
            if (!method.IsPageable)
            {
                throw new InvalidOperationException($"method {method.Owner}.{method.Name} is not a pageable operation");
            }

            CodeModel = method.CodeModel;
            ContentType = (CompositeTypeGo)method.ReturnType.Body;
            var returnSeq = ContentType.Properties.FirstOrDefault(p => p.ModelType is SequenceTypeGo);
            if (returnSeq == null)
            {
                throw new InvalidOperationException($"pageable return type {ContentType.Name} for method {method.Owner}.{method.Name} does not contain an array");
            }
            ElementType = returnSeq.ModelType.Cast<SequenceTypeGo>().ElementType;

            Documentation = $"Contains a page of {ReturnTypeName} values.";
            PreparerNeeded = !method.NextMethodExists(CodeModel.Methods.Cast<MethodGo>());

            var pageableExtension = method.Extensions[AzureExtensions.PageableExtension] as Newtonsoft.Json.Linq.JContainer;
            NextLink = CodeNamerGo.Instance.GetPropertyName((string)pageableExtension["nextLinkName"]);
            if (string.IsNullOrWhiteSpace(NextLink))
            {
                throw new InvalidOperationException($"method {method.Owner}.{method.Name} contains a null nextLinkName so it shouldn't be treated as a pageable operation");
            }
            ItemName = CodeNamerGo.Instance.GetPropertyName((string)pageableExtension["itemName"] ?? "value");

            IteratorType = new IteratorTypeGo(this);
            if (method.Deprecated)
            {
                DeprecationMessage = "The method for this type has been deprecated.";
            }
        }

        /// <summary>
        /// Gets the value of the x-ms-pageable:nextLinkName property.
        /// </summary>
        public string NextLink { get; }

        /// <summary>
        /// Gets the value of the x-ms-pageable:itemName property.
        /// </summary>
        public string ItemName { get; }

        /// <summary>
        /// Gets true if this response type needs a preparer to retrieve the next page of results.
        /// This is false if the swagger explicitly defines a next operation (i.e. x-ms-pageable:operationName).
        /// </summary>
        public bool PreparerNeeded { get; private set; }

        /// <summary>
        /// Gets the underlying type returned from the pageable operation (i.e. the wrapper around the array).
        /// </summary>
        public CompositeTypeGo ContentType { get; }

        /// <summary>
        /// Gets the element type, i.e. the type in the arrary.
        /// </summary>
        public IModelType ElementType { get; }

        /// <summary>
        /// Gets the page's return type name.
        /// </summary>
        public string ReturnTypeName => ElementType.HasInterface() ? ElementType.GetInterfaceName() : ElementType.Name.Value;

        /// <summary>
        /// Gets the name of the preparer method used to prepare the request for retrieving the next page of results.
        /// </summary>
        public string PreparerMethodName => $"{ContentType.Name.ToCamelCase()}Preparer";

        /// <summary>
        /// Gets the iterator type associated with this paged type.
        /// </summary>
        public IteratorTypeGo IteratorType { get; }

        /// <summary>
        /// Gets the name of the next results function field.
        /// </summary>
        public string FnFieldName => "fn";

        /// <summary>
        /// Gets the name of the results field.
        /// </summary>
        public string ResultFieldName => ContentType.Name.ToVariableName();

        /// <summary>
        /// The qualified name for the advancer method.
        /// </summary>
        public string AdvancerQualifiedName => $"{Name}.NextWithContext";

        /// <summary>
        /// Gets the signature of the function used to advance to the next page.
        /// </summary>
        public string NextPageFunctionSig => $"func(context.Context, {ContentType.Name}) ({ContentType.Name}, error)";

        public override string Fields()
        {
            return $"    {FnFieldName} {NextPageFunctionSig}\n    {ResultFieldName} {ContentType.Name}";
        }

        public override bool Equals(object other)
        {
            if (other == null)
            {
                return false;
            }

            if (ReferenceEquals(this, other))
            {
                return true;
            }

            if (other is PageTypeGo asMyType)
            {
                return string.Compare(Name, asMyType.Name, StringComparison.Ordinal) == 0;
            }

            return false;
        }

        public override int GetHashCode()
        {
            return Name.GetHashCode();
        }

        /// <summary>
        /// Updates the page with properties from the passed page. Throws an exception if pages are not equal.
        /// </summary>
        /// <param name="page"></param>
        public void MergeWithPage(PageTypeGo page)
        {
            if (!this.Equals(page))
            {
                throw new InvalidOperationException($"{nameof(MergeWithPage)} requires method to be a pageable operation");
            }
            PreparerNeeded |= page.PreparerNeeded;
        }
    }
}
